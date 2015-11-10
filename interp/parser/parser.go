package parser

import (
	"errors"
	"sync"
	"time"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/token"
	"github.com/mk2/yon/interp/word"
)

type parser struct {
	sync.Mutex
	state         stateFn
	input         kit.TokenScanner
	memo          kit.Memory
	words         chan kit.Word
	stoppedCh     kit.StoppedCh
	errorCh       kit.ErrorCh
	lastWord      kit.Word
	leftDelim     kit.Token
	rightDelim    kit.Token
	onceAgainWord bool
}

type stateFn func(*parser) stateFn

// New returns kit.Parser instance
func New(i kit.TokenScanner, memo kit.Memory) kit.Parser {

	p := &parser{
		state:         parse,
		input:         i,
		memo:          memo,
		words:         make(chan kit.Word),
		stoppedCh:     make(kit.StoppedCh),
		errorCh:       make(kit.ErrorCh),
		lastWord:      nil,
		leftDelim:     nil,
		rightDelim:    nil,
		onceAgainWord: false,
	}

	go p.run()

	return p
}

/*
================================================================================
Parser API
================================================================================
*/

func (p *parser) NextWord() kit.Word {

	word := <-p.words

	return word
}

func (p *parser) GetWords() <-chan kit.Word {

	return p.words
}

func (p *parser) ReadWord() (kit.Word, error) {

	if p.onceAgainWord {

		kit.Println("found unused last token")

		p.Lock()
		p.onceAgainWord = false
		p.Unlock()

		if p.lastWord == nil {
			return nil, errors.New("no last read token")
		}

		return p.lastWord, nil
	}

	kit.Println("waiting for incoming word")

	select {

	case t := <-p.words:
		p.lastWord = t
		return t, nil

	case <-time.After(kit.ParserTimeout):
		// timeout

	}

	return nil, errors.New("no word gained")
}

func (p *parser) UnreadWord() error {

	if p.onceAgainWord {
		return errors.New("already called UreadToken")
	}

	p.Lock()
	p.onceAgainWord = true
	p.Unlock()

	return nil
}

/*
================================================================================
parser private methods
================================================================================
*/

func (p *parser) run() {

	for p.state = parse; p.state != nil; {
		p.state = p.state(p)
	}
}

func (p *parser) emit(w kit.Word) {

	p.words <- w
}

func (p *parser) next() kit.Token {

	var (
		t   kit.Token
		err error
	)

	if t, err = p.input.ReadToken(); err != nil {
		return nil
	}

	return t
}

func (p *parser) peek() kit.Token {

	var (
		t   = p.next()
		err = p.input.UnreadToken()
	)

	if err != nil {
		return nil
	}

	return t
}

/*
================================================================================
parser functions
================================================================================
*/

func parse(p *parser) stateFn {

	p.leftDelim = nil
	p.rightDelim = nil

	switch t := p.peek(); t.GetType() {

	case token.TIdentifier:
		return parseIdentifier(p)

	case token.TLeftBrace:
		p.leftDelim = t
		return parseArray(p)

	case token.TNumber:
		w := word.NewNumberWord(t.GetVal())
		p.emit(w)
		p.next()

	case token.TString:
		w := word.NewStringWord(t.GetVal())
		p.emit(w)
		p.next()

	case token.TLeftSquareBracket:
		p.leftDelim = t
		return parseAnonFunc(p)

	case token.TSpace:
		p.next()

	case token.TEOF:
		p.emit(word.NewNilWord())
		return nil

	}

	return parse
}

func parseIdentifier(p *parser) stateFn {

	t := p.next()
	ident := t.GetVal()

	if w := p.memo.Vocab().Read(ident); w != nil {
		p.emit(w)
	} else {
		p.emit(word.NewNameWord(ident))
	}

	return parse
}

func parseArray(p *parser) stateFn {

	// enabled: emit stealing

	// skip the first leftside brace "{"
	p.next()

	p.emit(parseWordChain(p))

	return parse
}

// parseAnonFunc parses `[ number | string | ident | array | tuple ]`
func parseAnonFunc(p *parser) stateFn {

	// skip first double colon
	p.next()

	b := word.ChainWordFuncBody(parseWordChain(p))

	p.emit(word.NewFuncWord("", author.NewUserAuthor(), b))

	return parse
}

func parseWordChain(p *parser) kit.ChainWord {

	w := word.NewChainWord()

PARSE_WORD_CHAIN_LOOP:
	for {
		switch t := p.peek(); t.GetType() {

		case token.TNumber:
			w.Push(word.NewNumberWord(t.GetVal()))
			p.next()

		case token.TString:
			w.Push(word.NewStringWord(t.GetVal()))
			p.next()

		case token.TIdentifier:
			ident := t.GetVal()
			if v := p.memo.Vocab().Read(ident); v != nil {
				w.Push(v)
			} else {
				w.Push(word.NewNameWord(ident))
			}
			p.next()

		case token.TKeyword:
			// if a TKeyword token found, it indicates a begin of DictWord.

		case token.TLeftBrace:
			p.next()
			w.Push(parseWordChain(p))

		case token.TRightBrace:
			p.next()
			break PARSE_WORD_CHAIN_LOOP

		case token.TRightSquareBracket:
			p.next()
			break PARSE_WORD_CHAIN_LOOP

		case token.TEOF:
			break PARSE_WORD_CHAIN_LOOP

		case token.TSpace:
			p.next()
		}
	}

	return w
}

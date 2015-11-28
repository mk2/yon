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
		return parseArrayDict(p)

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

	default:
		p.next()

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

func parseArrayDict(p *parser) stateFn {

	if aw, dw := parseArrayDictBody(p); aw != nil {
		p.emit(aw)
	} else if dw != nil {
		p.emit(dw)
	}

	return parse
}

// parseAnonFunc parses `[ number | string | ident | array | tuple | func ]`
func parseAnonFunc(p *parser) stateFn {

	p.emit(parseAnonFuncBody(p))

	return parse
}

func parseArrayDictBody(p *parser) (kit.ArrayWord, kit.DictWord) {

	// skip the first bracket
	p.next()

	if w, t := parseChainWordBody(p, word.TArrayWord); t == word.TArrayWord {
		return word.NewArrayWordFromChainWord(w), nil
	} else if t == word.TDictWord {
		return nil, word.NewDictWordFromChainWord(w)
	}

	return nil, nil
}

func parseAnonFuncBody(p *parser) kit.FuncWord {

	// skip the first square bracket
	p.next()

	if w, t := parseChainWordBody(p, word.TFuncWord); t == word.TFuncWord {
		return word.NewFuncWordFromChainWord("", author.NewUserAuthor(), w)
	}

	return nil
}

func parseChainWordBody(p *parser, expectType kit.WordType) (kit.ChainWord, kit.WordType) {

	w := word.NewChainWord()
	actualType := expectType

PARSE_WORD_CHAIN_LOOP:
	for {
		switch t := p.peek(); {

		case w.Size() > 0 && t.GetType() == token.TDblColon:
			// TDblColon indicates the word is a dict word
			actualType = word.TDictWord
			p.next()

		case t.GetType() == token.TNumber:
			w.Push(word.NewNumberWord(t.GetVal()))
			p.next()

		case t.GetType() == token.TString:
			w.Push(word.NewStringWord(t.GetVal()))
			p.next()

		case t.GetType() == token.TIdentifier:
			ident := t.GetVal()
			if v := p.memo.Vocab().Read(ident); v != nil {
				w.Push(v)
			} else {
				w.Push(word.NewNameWord(ident))
			}
			p.next()

		case t.GetType() == token.TLeftBrace:
			if aw, dw := parseArrayDictBody(p); aw != nil {
				w.Push(aw)
			} else if dw != nil {
				w.Push(dw)
			}

		case t.GetType() == token.TLeftSquareBracket:
			w.Push(parseAnonFuncBody(p))

		case expectType == word.TArrayWord && t.GetType() == token.TRightBrace:
			p.next()
			break PARSE_WORD_CHAIN_LOOP

		case expectType == word.TFuncWord && t.GetType() == token.TRightSquareBracket:
			p.next()
			break PARSE_WORD_CHAIN_LOOP

		case t.GetType() == token.TEOF:
			break PARSE_WORD_CHAIN_LOOP

		case t.GetType() == token.TSpace:
			p.next()
		}
	}

	return w, actualType
}

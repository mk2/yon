package parser

import (
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/token"
	"github.com/mk2/yon/interp/word"
)

type parser struct {
	state         stateFn
	input         kit.TokenScanner
	words         chan kit.Word
	stoppedCh     kit.StoppedCh
	errorCh       kit.ErrorCh
	lastWord      kit.Word
	onceAgainWord bool
}

type stateFn func(*parser) stateFn

func New(i kit.TokenScanner) kit.Parser {

	p := &parser{
		state:         parse,
		input:         i,
		words:         make(chan kit.Word),
		stoppedCh:     make(kit.StoppedCh),
		errorCh:       make(kit.ErrorCh),
		lastWord:      nil,
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

/*
================================================================================
parser functions
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

func parse(p *parser) stateFn {

	switch t := p.peek(); t.GetType() {

	case token.TNumber:
		w := word.NewNumberWord(t.GetVal())
		p.emit(w)

	case token.TString:
		w := word.NewStringWord(t.GetVal())
		p.emit(w)

	}

	return parse
}

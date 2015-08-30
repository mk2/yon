package parser

import "github.com/mk2/yon/interp/kit"

type parser struct {
	l         kit.Lexer
	words     chan kit.Word
	stoppedCh kit.StoppedCh
	errorCh   kit.ErrorCh
}

type stateFn func(*parser) stateFn

func New(l kit.Lexer) kit.Parser {

	return &parser{
		l:         l,
		words:     make(chan kit.Word),
		stoppedCh: make(kit.StoppedCh),
		errorCh:   make(kit.ErrorCh),
	}
}

func (p *parser) NextWord() kit.Word {

	word := <-p.words

	return word
}

func (p *parser) GetWords() <-chan kit.Word {

	return p.words
}

func (p *parser) run() {

}

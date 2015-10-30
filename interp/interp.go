package interp

import (
	"log"

	"errors"

	"io"

	"github.com/mk2/yon/interp/history"
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/lexer"
	"github.com/mk2/yon/interp/memory"
	"github.com/mk2/yon/interp/parser"
	"github.com/mk2/yon/interp/stack"
	"github.com/mk2/yon/interp/vocabulary"
	"github.com/mk2/yon/interp/word"
)

type interp struct {
	source    string
	memo      kit.Memory
	stoppedCh kit.StoppedCh
	errorCh   kit.ErrorCh
}

/*
================================================================================
Interpreter APIs
================================================================================
*/

// New returns new interpeter object
func New() kit.Interpreter {

	interp := &interp{
		memo:      memory.New(stack.New(), vocabulary.New(), history.New()),
		stoppedCh: make(chan struct{}),
		errorCh:   make(chan error),
	}

	return interp
}

func (ip *interp) PrintStack() {

	ip.memo.Stack().Print()
}

func (ip *interp) PrintVocab() {

}

func (ip *interp) PrintHistory() {

}

func (ip *interp) EvalAndWait(runes io.RuneScanner) error {

	ip.Eval(runes)
	return ip.Wait()
}

func (ip *interp) Wait() error {

	select {

	case <-ip.stoppedCh:
		return nil

	case err := <-ip.errorCh:
		return err

	}
}

func (ip *interp) Eval(runes kit.RuneScanner) (kit.StoppedCh, kit.ErrorCh) {

	tokens := lexer.New(runes)
	words := parser.New(tokens)

	go ip.run(words)

	return ip.stoppedCh, ip.errorCh
}

/*
================================================================================
Interpreter private methods
================================================================================
*/

func (ip *interp) run(words kit.WordScanner) {

	m := ip.memo

	var (
		w   kit.Word
		err error
	)

	log.Println("start RUN_LOOP")

RUN_LOOP:
	for {

		if w, err = words.ReadWord(); err != nil {
			ip.errorCh <- err
			break RUN_LOOP
		}

		log.Printf("word: %t", w)

		switch w.GetWordType() {

		case word.TNumberWord:
			log.Println("number word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TStringWord:
			log.Println("string word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TNilWord:
			log.Println("nil word")
			break RUN_LOOP

		default:
			log.Println("unknown word: %+v", w)
			ip.errorCh <- errors.New("unknown word")
			break RUN_LOOP

		}

	}

	log.Println("exit RUN_LOOP")

	ip.stoppedCh <- struct{}{}
}

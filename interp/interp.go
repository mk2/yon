package interp

import (
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

	ip.memo.Vocab().Print()
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
	words := parser.New(tokens, ip.memo)

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

	kit.Println("start RUN_LOOP")

RUN_LOOP:
	for {

		if w, err = words.ReadWord(); err != nil {
			ip.errorCh <- err
			break RUN_LOOP
		}

		kit.Printf("word: %t", w)

		switch w.GetWordType() {

		case word.TNumberWord:
			kit.Println("number word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TStringWord:
			kit.Println("string word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TArrayWord:
			kit.Println("array word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TFuncWord:
			kit.Println("func word")
			if _, err := w.Do(m); err != nil {
				ip.errorCh <- err
				break
			}

		case word.TNilWord:
			kit.Println("nil word")
			break RUN_LOOP

		default:
			kit.Printf("unknown word: %t\n", w)
			ip.errorCh <- errors.New("unknown word")
			break RUN_LOOP

		}

	}

	kit.Println("exit RUN_LOOP")

	ip.stoppedCh <- struct{}{}
}

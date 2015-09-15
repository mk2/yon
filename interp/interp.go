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
	"github.com/mk2/yon/interp/token"
	"github.com/mk2/yon/interp/vocabulary"
	"github.com/mk2/yon/interp/word"
)

type Interpreter struct {
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
func New() (interp *Interpreter) {

	interp = &Interpreter{
		memo:      memory.New(stack.New(), vocabulary.New(), history.New()),
		stoppedCh: make(chan struct{}),
		errorCh:   make(chan error),
	}

	return
}

// PrintStack prints current stack contents
func (interp *Interpreter) PrintStack() {

	interp.memo.Stack().Print()
}

func (interp *Interpreter) EvalAndWait(r io.RuneScanner) error {

	interp.Eval(r)
	return interp.Wait()
}

func (interp *Interpreter) Wait() error {

	select {

	case <-interp.stoppedCh:
		return nil

	case err := <-interp.errorCh:
		return err

	}
}

func (interp *Interpreter) Eval(runes kit.RuneScanner) (kit.StoppedCh, kit.ErrorCh) {

	tokens := lexer.New(runes)
	words := parser.New(tokens)

	go interp.run(words)

	return interp.stoppedCh, interp.errorCh
}

/*
================================================================================
Interpreter private methods
================================================================================
*/

func (interp *Interpreter) run(words kit.WordScanner) {

	m := interp.memo

	var (
		w   kit.Word
		err error
	)

	for {

		if w, err = words.ReadWord(); err != nil {
			interp.errorCh <- err
			break
		}

		switch w.GetWordType() {

		case word.TNumberWord:
			log.Println("number word")
			if _, err := w.Do(m); err != nil {
				interp.errorCh <- err
				break
			}

		case word.TStringWord:
			log.Println("string word")
			if _, err := w.Do(m); err != nil {
				interp.errorCh <- err
				break
			}

		case word.TNilWord:
			log.Println("nil word")
			break

		default:
			log.Println("unknown word: %+v", w)
			interp.errorCh <- errors.New("unknown word")
			break

		}

	}

	interp.stoppedCh <- struct{}{}
}

/*
================================================================================
Interpreter parse methods
================================================================================
*/

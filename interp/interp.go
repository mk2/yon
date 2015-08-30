package interp

import (
	"log"

	"errors"

	"io"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/lexer"
	"github.com/mk2/yon/interp/memory"
	"github.com/mk2/yon/interp/stack"
	"github.com/mk2/yon/interp/vocabulary"
	"github.com/mk2/yon/interp/word"
)

type stoppedCh chan struct{}
type errorCh chan error

type Interpreter struct {
	source    string
	program   chan kit.Word
	memo      kit.Memory
	stoppedCh stoppedCh
	errorCh   errorCh
}

/*
================================================================================
Interpreter APIs
================================================================================
*/

// New returns new interpeter object
func New() (interp *Interpreter) {

	interp = &Interpreter{
		program:   make(chan kit.Word),
		memo:      memory.New(stack.New(), vocabulary.New()),
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

func (interp *Interpreter) Eval(r io.RuneScanner) (stoppedCh, errorCh) {

	tokens := lexer.New(r).GetTokenCh()

	go interp.run()

	go func() {
		for {

			var w kit.Word

			switch t := <-tokens; t.Typ {

			case lexer.TSpace:
				continue

			case lexer.TIdentifier:
				w = &word.Word{}
				w.SetWordType(word.TNilWord)

			case lexer.TNumber:
				w = word.NewNumberWord(t.Val)

			case lexer.TString:
				w = word.NewStringWord(t.Val)

			case lexer.TEOF:
				interp.stoppedCh <- struct{}{}
				return

			default:
				w = word.NewNilWord()

			}

			interp.program <- w
		}
	}()

	return interp.stoppedCh, interp.errorCh
}

/*
================================================================================
Interpreter private methods
================================================================================
*/

func (interp *Interpreter) run() {

	m := interp.memo

	for {

		switch w := <-interp.program; w.GetWordType() {

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

func (interp *Interpreter) parseIdentifier() {

}

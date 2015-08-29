package interp

import (
	"bytes"
	"log"

	"errors"

	"github.com/mk2/yon/interp/lexer"
	"github.com/mk2/yon/interp/stack"
	"github.com/mk2/yon/interp/word"
)

type stoppedCh chan struct{}
type errorCh chan error

type Interpreter struct {
	source    string
	stack     stack.Stack
	program   chan word.Word
	stoppedCh stoppedCh
	errorCh   errorCh
}

// New returns new interpeter object
func New() (interp *Interpreter) {

	interp = &Interpreter{
		program:   make(chan word.Word),
		stack:     *stack.New(),
		stoppedCh: make(chan struct{}),
		errorCh:   make(chan error),
	}

	go interp.run()

	return
}

func (interp *Interpreter) Wait() {

	<-interp.stoppedCh
}

func (interp *Interpreter) PrintStack() {

	stack.Print(&interp.stack)
}

func (interp *Interpreter) Eval(r *bytes.Buffer) (stoppedCh, errorCh) {

	l := lexer.New(r)
	tokens := l.GetTokenCh()

	go func() {
		for {

			var w word.Word

			switch t := <-tokens; t.Typ {

			case lexer.TSpace:
				continue

			case lexer.TIdentifier:
				w = &word.BaseWord{}
				w.SetWordType(word.NilWordType)

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

func (interp *Interpreter) run() {

	for {

		switch w := <-interp.program; w.GetWordType() {

		case word.NumberWordType:
			log.Println("number word")
			interp.stack.PushFront(w)

		case word.StringWordType:
			log.Println("string word")
			interp.stack.PushFront(w)

		case word.NilWordType:
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

package interp

import (
	"bytes"
	"strconv"
	"fmt"
	"github.com/mk2/yon/interp/word"
	"github.com/mk2/yon/interp/stack"
)

type Interpreter struct {
	source   string
	programs []word.Word
	stack    stack.Stack
}

// NewInterp returns new interpeter object
func New() (interp *Interpreter) {

	interp = new(Interpreter)
	interp = &Interpreter{
		programs: make([]word.Word, 0),
		stack: stack.New(),
	}

	return
}

func (interp *Interpreter) ParseAndExec(source string) (err error) {

	if err = interp.Parse(source); err != nil {
		return
	}

	if err = interp.Exec(); err != nil {
		return
	}

	return
}


// Exec is used for yon program execution
func (interp *Interpreter) Exec() (err error) {

	for _, w := range interp.programs {

		if err = w.Exec(); err != nil {
			return
		}
	}

	return
}

// Parse is used for parsing single line program
func (interp *Interpreter) Parse(source string) error {

	interp.source = source

	var (
		next int = -1
		w *word.BaseWord
		e error
	)

	for i, c := range source {

		if i <= next {
			continue
		}

		switch c {

		case ' ':
			continue

		case '(':


		case ':', '[', ']':

		case '.':
			if next, w, e = interp.readEmbed(i + 1); e != nil {
				return e
			}
			interp.programs = append(interp.programs, w)

		case '+', '-', '/', '%':

		case '"':
			if next, w, e = interp.readString(i + 1); e != nil {
				return e
			}
			interp.programs = append(interp.programs, w)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if next, w, e = interp.readNumber(i); e != nil {
				return e
			}
			interp.programs = append(interp.programs, w)

		default:
			if next, w, e = interp.readIdent(i); e != nil {
				return e
			}
			interp.programs = append(interp.programs, w)

		}

	}

	fmt.Printf("Programs: %t", interp.programs)

	return nil
}

func (interp *Interpreter) readNumber(from int) (next int, w *word.NumberWord, err error) {

	w = &word.NumberWord{}

	buf := bytes.NewBuffer([]byte{})

	NUM_LOOP:
	for i, c := range interp.source[from:] {

		switch c {

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf.WriteRune(c)
			next = from + i

		default:
			next = from + i
			break NUM_LOOP

		}

	}

	// convert string to float
	if s := buf.String(); s != "<nil>" {
		var f float64
		if f, err = strconv.ParseFloat(s, 64); err == nil {
			w.SetWordType(word.NumberWordType)
			w.Number = f
			return
		}
	}

	return
}

func (interp *Interpreter) readString(from int) (next int, w *word.StringWord, err error) {

	w = &word.BaseWord{
		wordType: word.NotWordType,
	}

	buf := bytes.NewBuffer([]byte{})

	STR_LOOP:
	for i, c := range interp.source[from:] {

		switch c {

		case '"':
			next = from + i
			break STR_LOOP

		default:
			buf.WriteRune(c)
			next = from + i

		}

	}

	// convert string to float
	if s := buf.String(); s != "<nil>" {
		w.SetWordType(word.StringWordType)
		w.String = s
		return
	}

	return
}

func (interp *Interpreter) readIdent(from int) (next int, w *word.StringWord, err error) {

	w = &word.FuncWord{}

	buf := bytes.NewBuffer([]byte{})

	IDENT_LOOP:
	for i, c := range interp.source[from:] {

		switch c {

		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'X', 'Y', 'Z',
			'?', '_', '-':
			buf.WriteRune(c)
			next = from + i

		default:
			next = from + i
			break IDENT_LOOP

		}

	}

	// convert string to float
	if s := buf.String(); s != "<nil>" {
		w.SetWordType(word.FuncWordType)
		w.String = s
		return
	}

	return
}

func (interp *Interpreter) readEmbed(from int) (next int, w *word.BaseWord, err error) {

	w = &word.BaseWord{
		wordType: word.NotWordType,
	}

	buf := bytes.NewBuffer([]byte{})

	STR_LOOP:
	for i, c := range interp.source[from:] {

		switch c {

		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'X', 'Y', 'Z',
			'?', '_', '-':
			buf.WriteRune(c)
			next = from + i

		default:
			next = from + i
			break STR_LOOP

		}

	}

	// convert string to float
	if s := buf.String(); s != "<nil>" {
		w.SetWordType(word.EmbedFuncWordType)

		switch s {

		case "", "s":

		}

		return
	}

	return
}

func (interp *Interpreter) readComment(from int) (next int, w *word.BaseWord, err error) {

	w = &word.BaseWord{
		wordType: word.NotWordType,
	}

	buf := bytes.NewBuffer([]byte{})
	nestCount := 0

	COMM_LOOP:
	for i, c := range interp.source[from:] {

		switch c {

		case '(':
			nestCount++

		case ')':
			nestCount--
			next = from + i
			if nestCount == 0 {
				break COMM_LOOP
			}

		default:
			next = from + i
			break COMM_LOOP

		}

	}

	// convert string to float
	if s := buf.String(); s != "<nil>" {
		w.SetWordType(word.EmbedFuncWordType)

		switch s {

		case "", "s":

		}

		return
	}

	return
}

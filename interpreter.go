package yon
import (
	"bytes"
	"strconv"
	"fmt"
	"container/list"
)

type Interpreter struct {
	source   string
	programs []*Word
	stack    *list.List
	ip       int
	np       int
}

// NewInterp is used for creating new interpreter
func NewInterpreter() (interp *Interpreter) {

	interp = new(Interpreter)
	interp.programs = make([]*Word, 0)
	interp.stack = list.New()
	interp.ip = 0
	interp.np = 1

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

	stack := interp.stack

	for _, w := range interp.programs {

		if err = w.Exec(stack); err != nil {
			return
		}
	}

	return
}

// Exec is used for execution single line program
func (interp *Interpreter) Parse(source string) error {

	interp.source = source

	var (
		next int = -1
		w *Word
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
			if next, w, e = interp.readStr(i + 1); e != nil {
				return e
			}
			interp.programs = append(interp.programs, w)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if next, w, e = interp.readNum(i); e != nil {
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

func (interp *Interpreter) readNum(from int) (next int, w *Word, err error) {

	w = &Word{
		wordType: NotWordType,
	}

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
			w.wordType = NumWordType
			w.num = f
			return
		}
	}

	return
}

func (interp *Interpreter) readStr(from int) (next int, w *Word, err error) {

	w = &Word{
		wordType: NotWordType,
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
		w.wordType = StrWordType
		w.str = s
		return
	}

	return
}

func (interp *Interpreter) readIdent(from int) (next int, w *Word, err error) {

	w = &Word{
		wordType: NotWordType,
	}

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
		w.wordType = IdentWordType
		w.str = s
		return
	}

	return
}

func (interp *Interpreter) readEmbed(from int) (next int, w *Word, err error) {

	w = &Word{
		wordType: NotWordType,
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
		w.wordType = EmbedWordType

		switch s {

		case "", "s":
			w.embedWordType = StackOpEmbedWordKind

		}

		return
	}

	return
}

func (interp *Interpreter) readComment(from int) (next int, w *Word, err error) {

	w = &Word{
		wordType: NotWordType,
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
		w.wordType = EmbedWordType

		switch s {

		case "", "s":
			w.embedWordType = StackOpEmbedWordKind

		}

		return
	}

	return
}

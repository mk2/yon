package stack

import (
	"bytes"
	"container/list"

	"fmt"

	"sync"

	"github.com/fatih/color"
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	fBase       = "|%5d| %s"
	fNumberWord = `<number> %f`
	fStringWord = `<string> "%s"`
	fNameWord   = `<name> %s`
	fArrayWord  = `<array> {%s}`
	fFuncWord   = `<func> [name:%s author:%s]`
	fSep        = "-"
)

type stack struct {
	sync.Mutex
	list.List
}

func New() kit.Stack {

	return &stack{
		List: *list.New(),
	}
}

func (s *stack) Print() {

	depth := 0

	for e := s.Front(); e != nil; e = e.Next() {

		if w, ok := e.Value.(kit.Word); ok {
			printWord(depth, w)
		}

		depth++
	}
}

func printWord(depth int, w kit.Word) {

	switch t := w.GetWordType(); {

	case t == word.TNumberWord:
		printNumberWord(depth, w.(kit.NumberWord))

	case t == word.TStringWord:
		printStringWord(depth, w.(kit.StringWord))

	case t == word.TNameWord:
		printNameWord(depth, w.(kit.NameWord))

	case t == word.TFuncWord:
		printFuncWord(depth, w.(kit.FuncWord))

	case t == word.TArrayWord:
		printArrayWord(depth, w.(kit.ArrayWord))
	}
}

func printNumberWord(depth int, w kit.NumberWord) {

	color.Magenta(fBase, depth, formatNumberWord(w))
}

func formatNumberWord(w kit.NumberWord) string {

	return fmt.Sprintf(fNumberWord, w.Number())
}

func printStringWord(depth int, w kit.StringWord) {

	color.Cyan(fBase, depth, formatStringWord(w))
}

func formatStringWord(w kit.StringWord) string {

	return fmt.Sprintf(fStringWord, w.String())
}

func printNameWord(depth int, w kit.NameWord) {

	color.Yellow(fBase, depth, formatNameWord(w))
}

func formatNameWord(w kit.NameWord) string {

	return fmt.Sprintf(fNameWord, w.Name())
}

func printFuncWord(depth int, w kit.FuncWord) {

	color.Magenta(fBase, depth, formatFuncWord(w))
}

func formatFuncWord(w kit.FuncWord) string {

	return fmt.Sprintf(fFuncWord, w.Name(), w.GetAuthorType())
}

func printArrayWord(depth int, w kit.ArrayWord) {

	color.Green(fBase, depth, formatArrayWord(w))
}

func formatArrayWord(w kit.ArrayWord) string {

	var buf *bytes.Buffer
	if buf = bytes.NewBufferString(""); buf == nil {
		return ""
	}

	isFirst := true

	for _, c := range w.Array() {

		if !isFirst {
			buf.WriteRune(',')
		}

		switch c.GetWordType() {

		case word.TNumberWord:
			buf.WriteString(formatNumberWord(c.(kit.NumberWord)))

		case word.TStringWord:
			buf.WriteString(formatStringWord(c.(kit.StringWord)))

		case word.TNameWord:
			buf.WriteString(formatNameWord(c.(kit.NameWord)))

		case word.TFuncWord:
			buf.WriteString(formatFuncWord(c.(kit.FuncWord)))

		case word.TArrayWord:
			buf.WriteString(formatArrayWord(c.(kit.ArrayWord)))

		}

		isFirst = false
	}

	return fmt.Sprintf(fArrayWord, buf.String())
}

func (s *stack) Push(v interface{}) *list.Element {

	s.Lock()
	e := s.PushFront(v)
	s.Unlock()

	return e
}

func (s *stack) Pop() *list.Element {

	s.Lock()
	e := s.Front()
	s.Remove(e)
	s.Unlock()

	return e
}

func (s *stack) Peek() *list.Element {

	e := s.Front()

	return e
}

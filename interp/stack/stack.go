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
		printNumberWord(depth, w.(*word.NumberWord))

	case t == word.TStringWord:
		printStringWord(depth, w.(*word.StringWord))

	case t == word.TNameWord:
		printNameWord(depth, w.(*word.NameWord))

	case t == word.TArrayWord:
		printArrayWord(depth, w.(*word.ArrayWord))
	}
}

func printNumberWord(depth int, w *word.NumberWord) {

	color.Magenta(fBase, depth, formatNumberWord(w))
}

func formatNumberWord(w *word.NumberWord) string {

	return fmt.Sprintf(fNumberWord, w.Number)
}

func printStringWord(depth int, w *word.StringWord) {

	color.Cyan(fBase, depth, formatStringWord(w))
}

func formatStringWord(w *word.StringWord) string {

	return fmt.Sprintf(fStringWord, w.String)
}

func printNameWord(depth int, w *word.NameWord) {

	color.Yellow(fBase, depth, formatNameWord(w))
}

func formatNameWord(w *word.NameWord) string {

	return fmt.Sprintf(fNameWord, w.Name)
}

func printArrayWord(depth int, w *word.ArrayWord) {

	color.Green(fBase, depth, formatArrayWord(w))
}

func formatArrayWord(w *word.ArrayWord) string {

	var buf *bytes.Buffer
	if buf = bytes.NewBufferString(""); buf == nil {
		return ""
	}

	isFirst := true

	for _, c := range w.Array {

		if !isFirst {
			buf.WriteRune(',')
		}

		switch c.GetWordType() {

		case word.TNumberWord:
			buf.WriteString(formatNumberWord(c.(*word.NumberWord)))

		case word.TStringWord:
			buf.WriteString(formatStringWord(c.(*word.StringWord)))

		case word.TNameWord:
			buf.WriteString(formatNameWord(c.(*word.NameWord)))

		case word.TArrayWord:
			buf.WriteString(formatArrayWord(c.(*word.ArrayWord)))

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

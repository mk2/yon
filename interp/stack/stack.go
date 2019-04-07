package stack

import (
	"bytes"
	"container/list"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	fBase = "|%5d| %s"
	fSep  = "-"
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

func (s *stack) Print() string {
	depth := 0
	var buf bytes.Buffer

	buf.WriteString("\n")
	for e := s.Front(); e != nil; e = e.Next() {

		if w, ok := e.Value.(kit.Word); ok {
			buf.WriteString(printWord(depth, w))
			buf.WriteString("\n")
		}

		depth++
	}

	return buf.String()
}

func printWord(depth int, w kit.Word) string {
	switch w.GetWordType() {

	case word.TNumberWord:
		return printNumberWord(depth, w.(kit.NumberWord))

	case word.TStringWord:
		return printStringWord(depth, w.(kit.StringWord))

	case word.TNameWord:
		return printNameWord(depth, w.(kit.NameWord))

	case word.TBoolWord:
		return printBoolWord(depth, w.(kit.BoolWord))

	case word.TFuncWord:
		return printFuncWord(depth, w.(kit.FuncWord))

	case word.TArrayWord:
		return printArrayWord(depth, w.(kit.ArrayWord))
	}

	return "<unknown word type>"
}

func printNumberWord(depth int, w kit.NumberWord) string {
	return Sprintf(Magenta(fBase), depth, w.Format())
}

func printStringWord(depth int, w kit.StringWord) string {
	return Sprintf(Cyan(fBase), depth, w.Format())
}

func printNameWord(depth int, w kit.NameWord) string {
	return Sprintf(Brown(fBase), depth, w.Format())
}

func printBoolWord(depth int, w kit.BoolWord) string {
	return Sprintf(Red(fBase), depth, w.Format())
}

func printFuncWord(depth int, w kit.FuncWord) string {
	return Sprintf(Gray(fBase), depth, w.Format())
}

func printArrayWord(depth int, w kit.ArrayWord) string {
	return Sprintf(Green(fBase), depth, w.Format())
}

func (s *stack) Push(v kit.Word) kit.Word {
	s.Lock()
	e := s.PushFront(v)
	s.Unlock()

	return e.Value.(kit.Word)
}

func (s *stack) Pop() kit.Word {
	s.Lock()
	e := s.Front()
	if e == nil {
		s.Unlock()
		return nil
	}
	s.Remove(e)
	s.Unlock()

	return e.Value.(kit.Word)
}

func (s *stack) Peek() kit.Word {
	e := s.Front()
	if e == nil {
		return nil
	}

	return e.Value.(kit.Word)
}

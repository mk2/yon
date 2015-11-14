package stack

import (
	"container/list"

	"sync"

	"github.com/fatih/color"
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

	switch w.GetWordType() {

	case word.TNumberWord:
		printNumberWord(depth, w.(kit.NumberWord))

	case word.TStringWord:
		printStringWord(depth, w.(kit.StringWord))

	case word.TNameWord:
		printNameWord(depth, w.(kit.NameWord))

	case word.TBoolWord:
		printBoolWord(depth, w.(kit.BoolWord))

	case word.TFuncWord:
		printFuncWord(depth, w.(kit.FuncWord))

	case word.TArrayWord:
		printArrayWord(depth, w.(kit.ArrayWord))
	}
}

func printNumberWord(depth int, w kit.NumberWord) {

	color.Magenta(fBase, depth, w.Format())
}

func printStringWord(depth int, w kit.StringWord) {

	color.Cyan(fBase, depth, w.Format())
}

func printNameWord(depth int, w kit.NameWord) {

	color.Yellow(fBase, depth, w.Format())
}

func printBoolWord(depth int, w kit.BoolWord) {

	color.Red(fBase, depth, w.Format())
}

func printFuncWord(depth int, w kit.FuncWord) {

	color.Magenta(fBase, depth, w.Format())
}

func printArrayWord(depth int, w kit.ArrayWord) {

	color.Green(fBase, depth, w.Format())
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

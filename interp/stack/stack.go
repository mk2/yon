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

	color.Magenta(fBase, depth, w.Format())
}

func printStringWord(depth int, w kit.StringWord) {

	color.Cyan(fBase, depth, w.Format())
}

func printNameWord(depth int, w kit.NameWord) {

	color.Yellow(fBase, depth, w.Format())
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
	// TODO add nil check
	e := s.Front()
	s.Remove(e)
	s.Unlock()

	return e.Value.(kit.Word)
}

func (s *stack) Peek() kit.Word {

	// TODO add nil check
	e := s.Front()

	return e.Value.(kit.Word)
}

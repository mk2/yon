package stack

import (
	"container/list"

	"fmt"

	"sync"

	"github.com/fatih/color"
	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	fBase       = "|%3d| %s"
	fNumberWord = `<number> %f`
	fStringWord = `<string> "%s"`
	fSep        = "-"
)

type Stack struct {
	kit.Stack
	sync.Mutex
	list.List
}

func New() *Stack {

	return &Stack{
		List: *list.New(),
	}
}

func (s *Stack) Print() {

	depth := 0

	for e := s.Front(); e != nil; e = e.Next() {

		if w, ok := e.Value.(kit.Word); ok {

			switch t := w.GetWordType(); {

			case t == word.TNumberWord:
				color.Magenta(fBase, depth, fmt.Sprintf(fNumberWord, 0))

			case t == word.TStringWord:
				color.Cyan(fBase, depth, fmt.Sprintf(fStringWord, "0"))

			}
		}

		depth += 1
	}
}

func (s *Stack) Push(v interface{}) *list.Element {

	s.Lock()
	e := s.PushFront(v)
	s.Unlock()

	return e
}

func (s *Stack) Pop() *list.Element {

	s.Lock()
	e := s.Front()
	s.Remove(e)
	s.Unlock()

	return e
}

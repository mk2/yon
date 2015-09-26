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

			switch t := w.GetWordType(); {

			case t == word.TNumberWord:
				color.Magenta(fBase, depth, fmt.Sprintf(fNumberWord, 0.0))

			case t == word.TStringWord:
				color.Cyan(fBase, depth, fmt.Sprintf(fStringWord, "0"))

			}
		}

		if f, ok := e.Value.(float64); ok {
			color.Magenta(fBase, depth, fmt.Sprintf(fNumberWord, f))
		}

		if str, ok := e.Value.(string); ok {
			color.Cyan(fBase, depth, fmt.Sprintf(fStringWord, str))
		}

		depth++
	}
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

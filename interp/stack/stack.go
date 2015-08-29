package stack

import (
	"container/list"

	"fmt"

	"github.com/fatih/color"
	"github.com/mk2/yon/interp/word"
)

const (
	fBase       = "|%3d| %s"
	fNumberWord = `<number> %f`
	fStringWord = `<string> "%s"`
	fSep        = "-"
)

type Stack struct {
	list.List
}

func New() *Stack {

	return &Stack{
		List: *list.New(),
	}
}

func Print(s *Stack) {

	depth := 0

	for e := s.Front(); e != nil; e = e.Next() {

		if w, ok := e.Value.(word.Word); ok {

			switch t := w.GetWordType(); {

			case t == word.NumberWordType:
				color.Magenta(fBase, depth, fmt.Sprintf(fNumberWord, 0))

			case t == word.StringWordType:
				color.Cyan(fBase, depth, fmt.Sprintf(fStringWord, "0"))

			}
		}

		depth += 1
	}
}

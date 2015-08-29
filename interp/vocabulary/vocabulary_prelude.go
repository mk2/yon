package vocabulary

import (
	"fmt"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	kStackPrint = ".s"
	kPopPrint   = "."
	kDup        = "dup"
	kDef        = "def"
)

func (v *Vocabulary) LoadPrelude() error {

	v.Write(kPopPrint, &word.FuncWord{
		Name: kDup,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			fmt.Printf("%v¥n", s.Pop().Value)
			return nil
		},
	})

	v.Write(kDup, &word.FuncWord{
		Name: kDup,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek().Value)
			return nil
		},
	})

	return nil
}

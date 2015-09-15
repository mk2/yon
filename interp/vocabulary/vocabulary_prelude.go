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
	kForceDef   = "def!"
	kAskDef     = "def?"
	kApply      = "apply"
	kCall       = "call"
	kPlus       = "+"
	kMinus      = "-"
	kMulti      = "*"
	kDiv        = "/"
	kRem        = "%"
)

func (v *vocabulary) LoadPrelude() error {

	v.OverWrite(kPopPrint, &word.FuncWord{
		Name: kPopPrint,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			fmt.Printf("%v¥n", s.Pop().Value)
			return nil
		},
	})

	v.OverWrite(kDup, &word.FuncWord{
		Name: kDup,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek().Value)
			return nil
		},
	})

	return nil
}

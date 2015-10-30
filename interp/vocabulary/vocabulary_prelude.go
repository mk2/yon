package vocabulary

import (
	"fmt"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	VStackPrint = ".s"
	VPopPrint   = "."
	VDup        = "dup"
	VDef        = "def"
	VForceDef   = "def!"
	VAskDef     = "def?"
	VApply      = "apply"
	VCall       = "call"
	VPlus       = "+"
	VMinus      = "-"
	VMulti      = "*"
	VDiv        = "/"
	VRem        = "%"
)

func (v *vocabulary) LoadPrelude() error {

	v.OverWrite(VPopPrint, word.NewFuncWord(
		VPopPrint,
		func(m kit.Memory) error {
			s := m.Stack()
			fmt.Printf("%v¥n", s.Pop().Value)
			return nil
		},
	))

	v.OverWrite(VDup, word.NewFuncWord(
		VDup,
		func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek().Value)
			return nil
		},
	))

	v.OverWrite(VDef, word.NewFuncWord(
		VDef,
		func(m kit.Memory) error {
			s := m.Stack()
			name := s.Pop().Value.(string)
			value := s.Pop().Value.(string)
			return v.Write(name, word.NewStringWord(value))
		},
	))

	return nil
}

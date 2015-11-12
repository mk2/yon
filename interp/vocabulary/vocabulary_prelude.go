package vocabulary

import (
	"errors"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	VStackPrint    = ".s"
	VPopPrint      = "."
	VDup           = "dup"
	VDef           = "def"
	VApply         = "apply"
	VCall          = "call"
	VEach          = "each"
	VMap           = "map"
	VRep           = "rep"
	VSh            = "sh"
	VPrint         = "print"
	VPrintSynonym  = "p"
	VPrintf        = "printf"
	VPrintfSynonym = "format"
	VPlus          = "+"
	VMinus         = "-"
	VMulti         = "*"
	VDiv           = "/"
	VRem           = "%"
)

func (v *vocabulary) LoadPrelude() error {

	v.OverWrite(VPopPrint, word.NewPreludeFuncWord(
		VPopPrint,
		func(m kit.Memory) error {
			s := m.Stack()
			m.Printf("%v\n", s.Pop())
			return nil
		},
	))

	v.OverWrite(VPrint, word.NewPreludeFuncWord(
		VPrint,
		func(m kit.Memory) error {
			s := m.Stack()
			m.Printf("%v\n", s.Peek())
			return nil
		},
	))

	v.AliasOverWrite(VPrint, VPrintSynonym)

	v.OverWrite(VDup, word.NewPreludeFuncWord(
		VDup,
		func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek())
			return nil
		},
	))

	v.OverWrite(VDef, word.NewPreludeFuncWord(
		VDef,
		func(m kit.Memory) error {
			var nw = m.Stack().Pop()

			value := m.Stack().Pop()
			name := ""

			switch nw.GetWordType() {

			case word.TStringWord:
				name = nw.(kit.StringWord).String()

			case word.TNameWord:
				name = nw.(kit.NameWord).Name()

			}

			return v.Write(name, value)
		},
	))

	v.OverWrite(VEach, word.NewPreludeFuncWord(
		VEach,
		func(m kit.Memory) error {

			var (
				fn = m.Stack().Pop()
				w  = m.Stack().Pop()
			)

			if fn.GetWordType() != word.TFuncWord || w.GetWordType() != word.TArrayWord {
				return errors.New("invalid word gain")
			}

			return nil
		},
	))

	return nil
}

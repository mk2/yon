package vocabulary

import (
	"fmt"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	VStackPrint    = ".s"
	VPopPrint      = "."
	VDup           = "dup"
	VDef           = "def"
	VForceDef      = "def!"
	VAskDef        = "def?"
	VApply         = "apply"
	VCall          = "call"
	VEach          = "each"
	VMap           = "map"
	VFMap          = "fmap"
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
			fmt.Printf("%v\n", s.Pop().Value)
			return nil
		},
	))

	v.OverWrite(VPrint, word.NewPreludeFuncWord(
		VPrint,
		func(m kit.Memory) error {
			s := m.Stack()
			fmt.Printf("%v\n", s.Peek().Value)
			return nil
		},
	))

	v.AliasOverWrite(VPrint, VPrintSynonym)

	v.OverWrite(VDup, word.NewPreludeFuncWord(
		VDup,
		func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek().Value)
			return nil
		},
	))

	v.OverWrite(VDef, word.NewPreludeFuncWord(
		VDef,
		func(m kit.Memory) error {
			s := m.Stack()
			var nw = s.Pop().Value.(kit.Word)
			value := s.Pop().Value.(kit.Word)

			name := ""
			switch nw.GetWordType() {

			case word.TStringWord:
				name = nw.(*word.StringWord).String

			case word.TNameWord:
				name = nw.(*word.NameWord).Name

			}

			return v.Write(name, value)
		},
	))

	return nil
}

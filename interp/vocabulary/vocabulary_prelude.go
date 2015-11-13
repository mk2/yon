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
	VRot           = "rot"
	VOver          = "over"
	VDef           = "def"
	VApply         = "apply"
	VCall          = "call"
	VEach          = "each"
	VIf            = "if"
	VMap           = "map"
	VTimes         = "times"
	VSh            = "sh"
	VAddEh         = "addeh"
	VPrint         = "print"
	VPrintSynonym  = "p"
	VPrintf        = "printf"
	VPrintfSynonym = "format"
	VPlus          = "+"
	VMinus         = "-"
	VMulti         = "*"
	VDiv           = "/"
	VRem           = "%"
	VNil           = "nil"
	VEq            = "="
	VNEq           = "/="
	VGt            = ">"
	VLt            = "<"
	VGte           = ">="
	VLte           = "<="
)

func (v *vocabulary) LoadPrelude() error {

	v.NewClass("prelude")

	v.OverWrite(CPrelude, VStackPrint, word.NewPreludeFuncWord(
		VStackPrint,
		func(m kit.Memory, args ...interface{}) error {
			m.Stack().Print()
			return nil
		},
	))

	v.OverWrite(CPrelude, VOver, word.NewPreludeFuncWord(
		VOver,
		func(m kit.Memory, args ...interface{}) error {
			upper := m.Stack().Pop()
			bottom := m.Stack().Pop()
			m.Stack().Push(bottom)
			m.Stack().Push(upper)
			m.Stack().Push(bottom)
			return nil
		},
	))

	v.OverWrite(CPrelude, VPopPrint, word.NewPreludeFuncWord(
		VPopPrint,
		func(m kit.Memory, args ...interface{}) error {
			s := m.Stack()
			m.Printf("%v\n", s.Pop())
			return nil
		},
	))

	v.OverWrite(CPrelude, VPrint, word.NewPreludeFuncWord(
		VPrint,
		func(m kit.Memory, args ...interface{}) error {
			s := m.Stack()
			m.Printf("%v\n", s.Peek())
			return nil
		},
	))

	v.AliasOverWrite(CPrelude, VPrint, VPrintSynonym)

	v.OverWrite(CPrelude, VDup, word.NewPreludeFuncWord(
		VDup,
		func(m kit.Memory, args ...interface{}) error {
			s := m.Stack()
			s.Push(s.Peek())
			return nil
		},
	))

	v.OverWrite(CPrelude, VDef, word.NewPreludeFuncWord(
		VDef,
		func(m kit.Memory, args ...interface{}) error {
			var nw = m.Stack().Pop()

			value := m.Stack().Pop()
			name := ""

			switch nw.GetWordType() {

			case word.TStringWord:
				name = nw.(kit.StringWord).String()

			case word.TNameWord:
				name = nw.(kit.NameWord).Name()

			}

			return v.Write(CUser, name, value)
		},
	))

	v.OverWrite(CPrelude, VEach, word.NewPreludeFuncWord(
		VEach,
		func(m kit.Memory, args ...interface{}) error {

			var (
				fn = m.Stack().Pop()
				w  = m.Stack().Pop()
			)

			if fn.GetWordType() != word.TFuncWord || !word.CheckChainWord(w) {
				return errors.New("invalid word gain")
			}

			w.(kit.ChainWord).Each(func(wd kit.Word) {
				m.Stack().Push(wd)
				fn.Do(m)
			})

			return nil
		},
	))

	return nil
}

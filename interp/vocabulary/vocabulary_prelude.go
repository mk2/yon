package vocabulary

import (
	"errors"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	VHistPrint     = ".h"
	VVocabPrint    = ".v"
	VStackPrint    = ".s"
	VPopPrint      = "."
	VDup           = "dup"
	VRot           = "rot"
	VOver          = "over"
	VDef           = "def"
	VApply         = "apply"
	VEach          = "each"
	VIf            = "if"
	VMap           = "map"
	VTimes         = "times"
	VSh            = "sh"
	VAddHandler    = "addh"
	VChangeStack   = "chst"
	VPrint         = "printt"
	VPrintSynonym  = "p"
	VPrintf        = "printf"
	VPrintfSynonym = "pf"
	VPlus          = "+"
	VMinus         = "-"
	VMulti         = "*"
	VDiv           = "/"
	VRem           = "%"
	VEq            = "="
	VNEq           = "/="
	VGt            = ">"
	VLt            = "<"
	VGte           = ">="
	VLte           = "<="
	VNil           = "nil"
	VTrue          = "true"
	VFalse         = "false"
)

func (v *vocabulary) LoadPrelude() error {

	v.NewClass("prelude")

	v.OverWrite(CPrelude, VNil, word.NewNilWord())

	v.OverWrite(CPrelude, VTrue, word.NewBoolWord(true))

	v.OverWrite(CPrelude, VFalse, word.NewBoolWord(false))

	v.OverWrite(CPrelude, VVocabPrint, word.NewPreludeFuncWord(
		VVocabPrint,
		func(m kit.Memory, args ...interface{}) error {
			m.Println(m.Vocab().Print())
			return nil
		},
	))

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

	v.OverWrite(CPrelude, VRot, word.NewPreludeFuncWord(
		VRot,
		func(m kit.Memory, args ...interface{}) error {
			var ws = make([]kit.Word, 0)
			for w := m.Stack().Pop(); w != nil; w = m.Stack().Pop() {
				ws = append(ws, w)
			}
			for _, w := range ws {
				m.Stack().Push(w)
			}
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

	v.OverWrite(CPrelude, VIf, word.NewPreludeFuncWord(
		VIf,
		func(m kit.Memory, args ...interface{}) error {

			var (
				ifFalseFn = m.Stack().Pop()
				ifTrueFn  = m.Stack().Pop()
				boolW     = m.Stack().Pop()
			)

			if ifFalseFn.GetWordType() != word.TFuncWord ||
				ifTrueFn.GetWordType() != word.TFuncWord ||
				boolW.GetWordType() != word.TBoolWord {
				return errors.New("invalid stack values")
			}

			if boolW.GetWordType() == word.TBoolWord {
				if boolW.(kit.BoolWord).Eval() {
					ifTrueFn.Do(m)
				} else {
					ifFalseFn.Do(m)
				}
			}

			return nil
		},
	))

	v.OverWrite(CPrelude, VApply, word.NewPreludeFuncWord(
		VApply,
		func(m kit.Memory, args ...interface{}) error {

			if w := m.Stack().Peek(); w != nil {
				w.Do(m)
			}

			return nil
		},
	))

	//
	// Comparator functions {{{
	//

	v.OverWrite(CPrelude, VEq, word.NewPreludeFuncWord(
		VEq,
		func(m kit.Memory, args ...interface{}) error {

			var (
				rhs = m.Stack().Pop()
				lhs = m.Stack().Pop()
			)

			if rhs == nil || lhs == nil {
				return errors.New("nil word given")
			}

			return nil
		},
	))

	//
	// }}} Comparator functions
	//

	//
	// arithmetic operators {{{
	//

	// `+` operator
	v.OverWrite(CPrelude, VPlus, word.NewPreludeFuncWord(
		VPlus,
		func(m kit.Memory, args ...interface{}) error {

			var (
				rhs = m.Stack().Pop()
				lhs = m.Stack().Pop()
			)

			if rhs == nil || lhs == nil {
				return errors.New("nil word given")
			}

			if rhs.GetWordType() != word.TNumberWord || lhs.GetWordType() != word.TNumberWord {
				return errors.New("not number word given")
			}

			s := m.Stack()
			s.Push(word.NewNumberWordFromFloat64(lhs.(kit.NumberWord).Number() + rhs.(kit.NumberWord).Number()))

			return nil
		},
	))

	// `-` operator
	v.OverWrite(CPrelude, VMinus, word.NewPreludeFuncWord(
		VMinus,
		func(m kit.Memory, args ...interface{}) error {

			var (
				rhs = m.Stack().Pop()
				lhs = m.Stack().Pop()
			)

			if rhs == nil || lhs == nil {
				return errors.New("nil word given")
			}

			if rhs.GetWordType() != word.TNumberWord || lhs.GetWordType() != word.TNumberWord {
				return errors.New("not number word given")
			}

			s := m.Stack()
			s.Push(word.NewNumberWordFromFloat64(lhs.(kit.NumberWord).Number() - rhs.(kit.NumberWord).Number()))

			return nil
		},
	))

	//
	// }}} arithmetic operators
	//

	return nil
}

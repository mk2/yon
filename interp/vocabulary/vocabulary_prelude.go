package vocabulary

import (
	"fmt"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/word"
)

const (
	stackPrint = ".s"
	popPrint   = "."
	dup        = "dup"
	def        = "def"
	forceDef   = "def!"
	askDef     = "def?"
	apply      = "apply"
	call       = "call"
	plus       = "+"
	minus      = "-"
	multi      = "*"
	div        = "/"
	rem        = "%"
)

func (v *vocabulary) LoadPrelude() error {

	v.OverWrite(popPrint, &word.FuncWord{
		Name: popPrint,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			fmt.Printf("%v¥n", s.Pop().Value)
			return nil
		},
	})

	v.OverWrite(dup, &word.FuncWord{
		Name: dup,
		Body: func(m kit.Memory) error {
			s := m.Stack()
			s.Push(s.Peek().Value)
			return nil
		},
	})

	return nil
}

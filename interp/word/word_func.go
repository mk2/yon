package word

import "github.com/mk2/yon/interp/memory"

type FuncBody func() error

type FuncWord struct {
	BaseWord
	Name string
	Body FuncBody
}

func (w *FuncWord) Read(m *memory.Memory) (interface{}, error) {

	return nil, w.Body()
}

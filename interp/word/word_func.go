package word

import "github.com/mk2/yon/interp/kit"

type FuncBody func() error

type FuncWord struct {
	Word
	Name string
	Body FuncBody
}

func (w *FuncWord) Read(m kit.Memory) (interface{}, error) {

	return nil, w.Body()
}

package word

import "github.com/mk2/yon/interp/kit"

type FuncWord struct {
	Word
	Name string
	Body kit.WordFuncBody
}

func (w *FuncWord) Do(m kit.Memory) (interface{}, error) {

	return nil, w.Body(m)
}

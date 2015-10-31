package word

import "github.com/mk2/yon/interp/kit"

type ArrayWord struct {
	Word
	Array []kit.Word
}

func NewArrayWord() *ArrayWord {

	return &ArrayWord{
		Word:  Word{wordType: TArrayWord},
		Array: []kit.Word{},
	}
}

func (w *ArrayWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *ArrayWord) Put(wd kit.Word) {

	w.Array = append(w.Array, wd)
}

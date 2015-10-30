package word

import "github.com/mk2/yon/interp/kit"

type ArrayWord struct {
	Word
	Array []Word
}

func NewArrayWord() *ArrayWord {

	return &ArrayWord{
		Word:  Word{wordType: TArrayWord},
		Array: []Word{},
	}
}

func (arr *ArrayWord) Do(m kit.Memory) (interface{}, error) {

	return nil, nil
}

func (arr *ArrayWord) Put(w Word) {

	arr.Array = append(arr.Array, w)
}

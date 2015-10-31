package word

import "github.com/mk2/yon/interp/kit"

type NilWord struct {
	Word
}

func NewNilWord() *NilWord {

	return &NilWord{
		Word: Word{
			wordType: TNilWord,
		},
	}
}

func (w *NilWord) Do(m kit.Memory) (interface{}, error) {

	// do nothing

	return nil, nil
}

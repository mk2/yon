package word

import "github.com/mk2/yon/interp/kit"

type NumberWord struct {
	Word
	Number float64
}

func NewNumberWord(val string) *NumberWord {

	return &NumberWord{
		Number: 0,
		Word: Word{
			wordType: TNumberWord,
		},
	}
}

func (w *NumberWord) Do(m kit.Memory) (interface{}, error) {

	return w.Number, nil
}

type StringWord struct {
	Word
	String string
}

func NewStringWord(val string) *StringWord {

	return &StringWord{
		String: val,
		Word: Word{
			wordType: TStringWord,
		},
	}
}

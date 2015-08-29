package word

import "github.com/mk2/yon/interp/kit"

type NumberWord struct {
	BaseWord
	Number float64
}

func NewNumberWord(val string) *NumberWord {

	return &NumberWord{
		Number: 0,
		BaseWord: BaseWord{
			wordType: TNumberWord,
		},
	}
}

func (w *NumberWord) Read(m kit.Memory) (interface{}, error) {

	return w.Number, nil
}

type StringWord struct {
	BaseWord
	String string
}

func NewStringWord(val string) *StringWord {

	return &StringWord{
		String: val,
		BaseWord: BaseWord{
			wordType: TStringWord,
		},
	}
}

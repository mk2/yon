package word

import (
	"strconv"

	"github.com/mk2/yon/interp/kit"
)

type NumberWord struct {
	Word
	Number float64
}

func NewNumberWord(val string) *NumberWord {

	var (
		f   float64
		err error
	)
	if f, err = strconv.ParseFloat(val, 64); err != nil {
		f = 0.0
	}

	return &NumberWord{
		Number: f,
		Word: Word{
			wordType: TNumberWord,
		},
	}
}

func (w *NumberWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w.Number)

	return nil, nil
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

func (w *StringWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w.String)

	return nil, nil
}

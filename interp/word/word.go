package word

import "github.com/mk2/yon/interp/kit"

const (
	TNilWord kit.WordType = iota
	TFuncWord
	TNumberWord
	TStringWord
	TArrayWord
)

type Word struct {
	wordType kit.WordType
}

func (w *Word) GetWordType() kit.WordType {

	return w.wordType
}

func (w *Word) SetWordType(wordType kit.WordType) {

	w.wordType = wordType
}

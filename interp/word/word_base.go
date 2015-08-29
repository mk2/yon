package word

import "github.com/mk2/yon/interp/kit"

type Word struct {
	kit.Word
	wordType kit.WordType
}

func (w *Word) GetWordType() kit.WordType {

	return w.wordType
}

func (w *Word) SetWordType(wordType kit.WordType) {

	w.wordType = wordType
}

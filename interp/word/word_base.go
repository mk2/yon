package word

import "github.com/mk2/yon/interp/kit"

type BaseWord struct {
	kit.Word
	wordType kit.WordType
}

func (w *BaseWord) GetWordType() kit.WordType {

	return w.wordType
}

func (w *BaseWord) SetWordType(wordType kit.WordType) {

	w.wordType = wordType
}

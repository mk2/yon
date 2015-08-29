package word

type BaseWord struct {
	Word
	wordType WordType
}

func (w *BaseWord) GetWordType() WordType {

	return w.wordType
}

func (w *BaseWord) SetWordType(wordType WordType) {

	w.wordType = wordType
}

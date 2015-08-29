package word

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

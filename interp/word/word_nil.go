package word

type NilWord struct {
	BaseWord
}

func NewNilWord() *NilWord {

	return &NilWord{
		BaseWord: BaseWord{
			wordType: NilWordType,
		},
	}
}

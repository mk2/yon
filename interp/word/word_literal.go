package word

type NumberWord struct {
	BaseWord
	Number float64
}

func NewNumberWord(val string) *NumberWord {

	return &NumberWord{
		Number: 0,
		BaseWord: BaseWord{
			wordType: NumberWordType,
		},
	}
}

type StringWord struct {
	BaseWord
	String string
}

func NewStringWord(val string) *StringWord {

	return &StringWord{
		String: val,
		BaseWord: BaseWord{
			wordType: StringWordType,
		},
	}
}

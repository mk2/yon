package word

import "github.com/mk2/yon/interp/kit"

type FuncWord struct {
	Word
	Name string
	Body kit.WordFuncBody
}

func NewPreludeFuncWord(name string, body kit.WordFuncBody) *FuncWord {

	return &FuncWord{
		Word: Word{wordType: TFuncWord, authorType: AuthorPrelude},
		Name: name,
		Body: body,
	}
}

func NewFuncWord(name string, author kit.AuthorType, body kit.WordFuncBody) *FuncWord {

	return &FuncWord{
		Word: Word{wordType: TFuncWord, authorType: author},
		Name: name,
		Body: body,
	}
}

func (w *FuncWord) Do(m kit.Memory) (interface{}, error) {

	return nil, w.Body(m)
}

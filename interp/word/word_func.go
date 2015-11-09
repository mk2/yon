package word

import (
	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type funcWord struct {
	word
	name string
	body kit.WordFuncBody
}

func NewPreludeFuncWord(name string, body kit.WordFuncBody) kit.FuncWord {

	return NewFuncWord(name, author.NewPreludeAuthor(), body)
}

func NewFuncWord(name string, author kit.Author, body kit.WordFuncBody) kit.FuncWord {

	return &funcWord{
		word: word{wordType: TFuncWord, author: author},
		name: name,
		body: body,
	}
}

func (w *funcWord) Do(m kit.Memory) (interface{}, error) {

	return nil, w.body(m)
}

func ArrayWordFuncBody(a *arrayWord) kit.WordFuncBody {

	return func(m kit.Memory) (err error) {

		for _, w := range a.array {
			_, err = w.Do(m)
		}

		return err
	}
}

func (f *funcWord) String() string {

	return ""
}

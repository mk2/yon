package word

import (
	"container/list"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type funcWord struct {
	chainWord
	name string
	body kit.WordFuncBody
}

func NewPreludeFuncWord(name string, body kit.WordFuncBody) kit.FuncWord {

	return NewFuncWord(name, author.NewPreludeAuthor(), body)
}

func NewFuncWord(name string, author kit.Author, body kit.WordFuncBody) kit.FuncWord {

	return &funcWord{
		chainWord: chainWord{
			word: word{wordType: TFuncWord, author: author},
			List: *list.New(),
		},
		name: name,
		body: body,
	}
}

func NewFuncWordFromChainWord(name string, author kit.Author, c kit.ChainWord) kit.FuncWord {

	return NewFuncWord(name, author, funcBodyFromChainWord(c))
}

func (w *funcWord) Do(m kit.Memory) (interface{}, error) {

	return nil, w.body(m)
}

func funcBodyFromChainWord(a kit.ChainWord) kit.WordFuncBody {

	return func(m kit.Memory) (err error) {

		for _, w := range a.(kit.ArrayWord).Array() {
			_, err = w.Do(m)
		}

		return err
	}
}

func (w *funcWord) Name() string {

	return w.name
}

func (f *funcWord) String() string {

	return ""
}

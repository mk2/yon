package word

import (
	"container/list"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type arrayWord struct {
	chainWord
	array []kit.Word
}

func NewArrayWordFromChainWord(c kit.ChainWord) kit.ArrayWord {

	return NewArrayWordFromList(c.ExtractList())
}

func NewArrayWord() kit.ArrayWord {

	return NewArrayWordFromList(list.New())
}

func NewArrayWordFromList(l *list.List) kit.ArrayWord {

	return &arrayWord{
		chainWord: chainWord{
			word: word{wordType: TArrayWord, author: author.NewUserAuthor()},
			List: *l,
		},
		array: []kit.Word{},
	}
}

func (w *arrayWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *arrayWord) Put(wd kit.Word) {

	w.array = append(w.array, wd)
}

func (w *arrayWord) Array() []kit.Word {

	return w.array
}

func (w *arrayWord) String() string {

	return ""
}

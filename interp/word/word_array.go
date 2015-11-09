package word

import (
	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type arrayWord struct {
	word
	array []kit.Word
}

func NewArrayWord() kit.ArrayWord {

	return &arrayWord{
		word:  word{wordType: TArrayWord, author: author.NewUserAuthor()},
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

func (w *arrayWord) String() string {

	return ""
}

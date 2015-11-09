package word

import (
	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type nameWord struct {
	word
	name string
}

func NewNameWord(name string) kit.NameWord {

	return &nameWord{
		word: word{wordType: TNameWord, author: author.NewUserAuthor()},
		name: name,
	}
}
func (w *nameWord) Do(m kit.Memory) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *nameWord) Name() string {

	return w.name
}

func (w *nameWord) String() string {

	return w.name
}

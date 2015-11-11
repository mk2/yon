package word

import (
	"container/list"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type chainWord struct {
	word
	list.List
}

func NewChainWord() kit.ChainWord {

	return &chainWord{
		word: word{wordType: TChainWord, author: author.NewUserAuthor()},
		List: *list.New(),
	}
}

func (w *chainWord) ExtractList() *list.List {

	return &w.List
}

func (w *chainWord) String() string {

	return ""
}

func (w *chainWord) Do(m kit.Memory) (interface{}, error) {

	return nil, nil
}

func (w *chainWord) Push(v kit.Word) kit.Word {

	return w.PushFront(v).Value.(kit.Word)
}

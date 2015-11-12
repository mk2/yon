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
	}
}

func (w *chainWord) ExtractList() *list.List {

	return &w.List
}

func (w *chainWord) String() string {

	return "chain"
}

func (w *chainWord) Format() string {

	return fChainWord
}

func (w *chainWord) Do(m kit.Memory, args ...interface{}) (interface{}, error) {

	return nil, nil
}

func (w *chainWord) Push(v kit.Word) kit.Word {

	return w.PushBack(v).Value.(kit.Word)
}

func (w *chainWord) Each(f func(kit.Word)) {

	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value != nil {
			w := e.Value.(kit.Word)
			f(w)
		}
	}
}

func (w *chainWord) FlatEach(f func(kit.Word)) {

	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value != nil {
			w := e.Value.(kit.Word)
			switch w.GetWordType() {
			case TChainWord:
				w.(kit.ChainWord).FlatEach(f)
			default:
				f(w)
			}
		}
	}
}

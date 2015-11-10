package word

import (
	"container/list"
	"sync"

	"github.com/mk2/yon/interp/kit"
)

type chainWord struct {
	word
	sync.Mutex
	list.List
}

func NewChainWord() kit.ChainWord {

	return nil
}

func (w *chainWord) Push(v kit.Word) kit.Word {

	return nil
}

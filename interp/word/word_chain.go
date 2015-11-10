package word

import (
	"container/ring"

	"github.com/mk2/yon/interp/kit"
)

type chainWord struct {
	r ring.Ring
}

func NewChainWord() kit.ChainWord {

	return nil
}

func wordToRing(w kit.Word) *ring.Ring {

	r := ring.New(1)
	r.Value = w

	return r
}

func (w *chainWord) Push(v kit.Word) kit.Word {

	return w.r.Link(wordToRing(v)).Value.(kit.Word)
}
func (w *chainWord) Pop() kit.Word {

	return nil
}

func (w *chainWord) Peek() kit.Word {

	return nil
}

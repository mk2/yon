package vocabulary

import (
	"sync"

	"github.com/mk2/yon/interp/kit"
)

type Vocabulary struct {
	kit.Vocabulary
	sync.Mutex
	words map[string]kit.Word
}

func New() (v *Vocabulary) {

	v = &Vocabulary{
		words: make(map[string]kit.Word, 0),
	}

	v.loadPrelude()

	return
}

func (v *Vocabulary) Write(k string, w kit.Word) {

	v.words[k] = w
}

func (v *Vocabulary) Read(k string) kit.Word {

	return v.words[k]
}

package vocabulary

import (
	"sync"

	"errors"

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

	v.LoadPrelude()

	return
}

func (v *Vocabulary) Write(k string, w kit.Word) error {

	if _, ok := v.words[k]; ok {
		return errors.New("already exists key: " + k)
	}

	v.words[k] = w

	return nil
}

func (v *Vocabulary) Read(k string) kit.Word {

	return v.words[k]
}

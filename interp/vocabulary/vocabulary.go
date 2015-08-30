package vocabulary

import (
	"sync"

	"errors"

	"github.com/mk2/yon/interp/kit"
)

type vocabulary struct {
	sync.Mutex
	words map[string]kit.Word
}

func New() kit.Vocabulary {

	v := &vocabulary{
		words: make(map[string]kit.Word, 0),
	}

	v.LoadPrelude()

	return v
}

func (v *vocabulary) Write(k string, w kit.Word) error {

	if _, ok := v.words[k]; ok {
		return errors.New("already exists key: " + k)
	}

	v.words[k] = w

	return nil
}

func (v *vocabulary) Read(k string) kit.Word {

	return v.words[k]
}

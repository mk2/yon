package vocabulary

import (
	"log"
	"sync"

	"errors"

	"github.com/mk2/yon/interp/kit"
)

type vocabulary struct {
	sync.Mutex
	sync.Once
	words map[string]kit.Word
}

func New() kit.Vocabulary {

	v := &vocabulary{
		words: make(map[string]kit.Word, 0),
	}

	v.LoadPrelude()

	return v
}

func (v *vocabulary) Print() {

	for k, w := range v.words {
		log.Printf("key: %t body:%t", k, w)
	}
}

func (v *vocabulary) Write(k string, w kit.Word) error {

	if _, ok := v.words[k]; ok {
		return errors.New("already exists key: " + k)
	}

	v.Lock()
	v.words[k] = w
	v.Unlock()

	return nil
}

func (v *vocabulary) OverWrite(k string, w kit.Word) (err error) {

	if _, ok := v.words[k]; ok {
		err = errors.New("already exists key: " + k)
	}

	v.Lock()
	v.words[k] = w
	v.Unlock()

	return err
}

func (v *vocabulary) AliasOverWrite(orig string, alter string) (err error) {

	if w, ok := v.words[orig]; ok {
		return v.OverWrite(alter, w)
	}

	return errors.New("not found " + orig)
}

func (v *vocabulary) Read(k string) kit.Word {

	return v.words[k]
}

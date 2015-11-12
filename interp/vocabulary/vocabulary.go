package vocabulary

import (
	"log"
	"sync"

	"errors"

	"github.com/mk2/yon/interp/kit"
	"strings"
)

type vocabulary struct {
	sync.Mutex
	sync.Once
	words   map[string]kit.Word
	classes map[string]map[string]kit.Word
}

func New() kit.Vocabulary {

	v := &vocabulary{
		words:   make(map[string]kit.Word, 0),
		classes: make(map[string]map[string]kit.Word, 0),
	}

	v.LoadPrelude()

	return v
}

func (v *vocabulary) NewClass(className string) error {

	v.Lock()
	v.classes[className] = make(map[string]kit.Word, 0)
	v.Unlock()

	return nil
}

func (v *vocabulary) Print() {

	for k, w := range v.words {
		log.Printf("key: %v body:%+v", k, w)
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

func (v *vocabulary) ReadClass(className string, k string) kit.Word {

	if class, classOk := v.classes[className]; classOk {
		if w, wordOk := class[k]; wordOk {
			return w
		}
	}

	return nil
}

func (v *vocabulary) Read(fqk string) kit.Word {

	return v.words[fqk]
}

// ExtractClass returns the extracted class name and key from the given fully qualified key.
func ExtractClass(fqk string) (string, string) {

	fqkLen := len(fqk)
	names := strings.Split(fqk, ".")
	key := names[len(names)-1]
	classEnd := len(fqk)-len(key)-1

	if classEnd < 0 {
		return "", key
	} else {
		return fqk[:fqkLen - len(key) - 1], key
	}
}

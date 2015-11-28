package vocabulary

import (
	"sync"

	"errors"

	"strings"

	"bytes"
	"fmt"

	"github.com/mk2/yon/interp/kit"
)

type vocabulary struct {
	sync.Mutex
	sync.Once
	classes map[string]map[string]kit.Word
}

const (
	CPrelude = "prelude"
	CPsUtil  = "psutil"
	CUser    = "user"
)

const (
	ClassSep = "~"
)

func New() kit.Vocabulary {

	v := &vocabulary{
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

func (v *vocabulary) Print() string {

	buf := new(bytes.Buffer)

	for c, ws := range v.classes {
		for k, w := range ws {
			fmt.Fprintf(buf, "%-16s :: %s\n", (c + ClassSep + k), w.Format())
		}
	}

	return buf.String()
}

func (v *vocabulary) ExistWord(c string, k string) (kit.Word, error) {

	if _, classOk := v.classes[c]; !classOk {
		return nil, errors.New("class not found: " + c)
	}

	if _, wordOk := v.classes[c][k]; !wordOk {
		return nil, errors.New("word not found: " + k)
	}

	return v.classes[c][k], nil
}

func (v *vocabulary) Write(c string, k string, w kit.Word) error {

	if _, classOk := v.classes[c]; !classOk {
		return errors.New("class not found: " + c)
	}

	if _, wordOk := v.classes[c][k]; wordOk {
		return errors.New("already exists key: " + k)
	}

	v.Lock()
	v.classes[c][k] = w
	v.Unlock()

	return nil
}

func (v *vocabulary) OverWrite(c string, k string, w kit.Word) (err error) {

	err = v.Write(c, k, w)

	if err != nil {
		v.Lock()
		v.classes[c][k] = w
		v.Unlock()
	}

	return err
}

func (v *vocabulary) AliasOverWrite(c string, k string, alter string) (err error) {

	if w := v.ReadClass(c, k); w != nil {
		v.OverWrite(c, alter, w)
	}

	return errors.New("not found: " + c + "~" + k)
}

func (v *vocabulary) Nil() kit.Word {

	return v.ReadClass(CPrelude, VNil)
}

func (v *vocabulary) True() kit.Word {

	return v.ReadClass(CPrelude, VTrue)
}

func (v *vocabulary) False() kit.Word {

	return v.ReadClass(CPrelude, VFalse)
}

func (v *vocabulary) ReadClass(c string, k string) kit.Word {

	if class, classOk := v.classes[c]; classOk {
		if w, wordOk := class[k]; wordOk {
			return w
		}
	}

	return nil
}

func (v *vocabulary) Read(fqk string) kit.Word {

	c, k := ExtractClass(fqk)

	if c == "" {
		if class, preludeOk := v.classes[CPrelude]; preludeOk {
			if w, classOk := class[k]; classOk {
				return w
			}
		}
		if class, psutilOk := v.classes[CPsUtil]; psutilOk {
			if w, classOk := class[k]; classOk {
				return w
			}
		}
		if class, userOk := v.classes[CUser]; userOk {
			if w, classOk := class[k]; classOk {
				return w
			}
		}
	}

	return v.ReadClass(c, k)
}

// ExtractClass returns the extracted class name and key from the given fully qualified key.
func ExtractClass(fqk string) (string, string) {

	fqkLen := len(fqk)
	names := strings.Split(fqk, ClassSep)
	key := names[len(names)-1]
	classEnd := len(fqk) - len(key) - 1

	if classEnd < 0 {
		return "", key
	} else {
		return fqk[:fqkLen-len(key)-1], key
	}
}

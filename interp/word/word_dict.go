package word

import "github.com/mk2/yon/interp/kit"

type dictWord struct {
	word
	d map[kit.Word]kit.Word
}

func NewDictWord() kit.DictWord {

	return nil
}

func (w *dictWord) Put(k kit.Word, v kit.Word) {

	w.d[k] = v
}

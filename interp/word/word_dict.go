package word

import (
	"bytes"
	"container/list"
	"fmt"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type dictWord struct {
	dict map[kit.Word]kit.Word
	chainWord
}

func NewDictWord() kit.DictWord {

	return NewDictWordFromList(list.New())
}

func NewDictWordFromChainWord(c kit.ChainWord) kit.DictWord {

	return NewDictWordFromList(c.ExtractList())
}

func NewDictWordFromList(l *list.List) kit.DictWord {

	return &dictWord{
		dict: make(map[kit.Word]kit.Word, 0),
		chainWord: chainWord{
			word: word{wordType: TArrayWord, author: author.NewUserAuthor()},
			List: *l,
		},
	}
}

func (w *dictWord) Map() map[kit.Word]kit.Word {

	return w.dict
}

func (w *dictWord) Do(m kit.Memory, args ...interface{}) (interface{}, error) {

	m.Stack().Push(w)

	return nil, nil
}

func (w *dictWord) Put(k kit.Word, v kit.Word) {

	w.dict[k] = v

	return
}

func (w *dictWord) String() string {

	var (
		buf     = new(bytes.Buffer)
		isFirst = true
	)
	w.Each(func(wd kit.Word) {
		if !isFirst {
			buf.WriteString(", ")
		}
		buf.WriteString(wd.String())
		isFirst = false
	})
	return buf.String()
}

func (w *dictWord) Format() string {

	var (
		buf     = new(bytes.Buffer)
		isFirst = true
	)
	w.Each(func(wd kit.Word) {
		if !isFirst {
			buf.WriteString(", ")
		}
		buf.WriteString(wd.Format())
		isFirst = false
	})
	return fmt.Sprintf(fDictWord, buf.String())
}

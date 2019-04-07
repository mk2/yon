package word

import (
	"bytes"
	"container/list"
	"fmt"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type arrayWord struct {
	chainWord
}

func NewArrayWordFromChainWord(c kit.ChainWord) kit.ArrayWord {
	return NewArrayWordFromList(c.ExtractList())
}

func NewArrayWord() kit.ArrayWord {
	return NewArrayWordFromList(list.New())
}

func NewArrayWordFromList(l *list.List) kit.ArrayWord {
	return &arrayWord{
		chainWord: chainWord{
			word: word{wordType: TArrayWord, author: author.NewUserAuthor()},
			List: *l,
		},
	}
}

func (w *arrayWord) Do(m kit.Memory, args ...interface{}) (interface{}, error) {
	m.Stack().Push(w)

	return nil, nil
}

func (w *arrayWord) Put(wd kit.Word) {
	w.PushBack(wd)
}

func (w *arrayWord) Array() []kit.Word {
	var ws []kit.Word
	w.Each(func(wd kit.Word) {
		ws = append(ws, wd)
	})

	return ws
}

func (w *arrayWord) String() string {
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

func (w *arrayWord) Format() string {
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
	return fmt.Sprintf(fArrayWord, buf.String())
}

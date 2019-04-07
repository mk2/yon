package word

import (
	"fmt"

	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type funcWord struct {
	chainWord
	name   string
	quoted bool
	body   kit.WordFuncBody
}

func NewPreludeFuncWord(name string, body kit.WordFuncBody) kit.FuncWord {
	return NewFuncWord(name, author.NewPreludeAuthor(), body)
}

func NewFuncWord(name string, author kit.Author, body kit.WordFuncBody) kit.FuncWord {
	return &funcWord{
		chainWord: chainWord{
			word: word{wordType: TFuncWord, author: author},
		},
		name:   name,
		body:   body,
		quoted: name == "",
	}
}

func NewFuncWordFromChainWord(name string, author kit.Author, c kit.ChainWord) kit.FuncWord {
	return NewFuncWord(name, author, funcBodyFromChainWord(c))
}

func (w *funcWord) Do(m kit.Memory, args ...interface{}) (interface{}, error) {
	if w.quoted {
		w.quoted = false
		m.Stack().Push(w)
		return nil, nil
	} else {
		return nil, w.body(m, args...)
	}
}

func funcBodyFromChainWord(a kit.ChainWord) kit.WordFuncBody {
	return func(m kit.Memory, args ...interface{}) (err error) {
		a.Each(func(w kit.Word) {
			_, err = w.Do(m)
		})

		return err
	}
}

func (w *funcWord) Name() string {
	return w.name
}

func (w *funcWord) String() string {
	return w.name
}

func (w *funcWord) Format() string {
	return fmt.Sprintf(fFuncWord, w.name, w.quoted)
}

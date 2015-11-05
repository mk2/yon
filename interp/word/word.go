package word

import (
	"github.com/mk2/yon/interp/kit"
	"sync"
)

const (
	TNilWord kit.WordType = iota
	TFuncWord
	TNumberWord
	TStringWord
	TArrayWord
	TNameWord
)

const (
	fNumberWord = `<number> %f`
	fStringWord = `<string> "%s"`
	fNameWord   = `<name> %s`
	fArrayWord  = `<array> {%s}`
	fFuncWord   = `<func> [name:%s authorType:%s authorId:%s]`
)

const (
	AuthorPrelude kit.AuthorType = "prelude"
	AuthorUser    kit.AuthorType = "user"
)

type Word struct {
	sync.Once
	wordType   kit.WordType
	authorType kit.AuthorType
	authorId   kit.AuthorId
}

func (w *Word) GetWordType() kit.WordType {

	return w.wordType
}

func (w *Word) SetWordType(wordType kit.WordType) {

	w.wordType = wordType
}

func (w *Word) GetAuthorType() kit.AuthorType {

	return w.authorType
}

func (w *Word) SetAuthorType(authorType kit.AuthorType) {

	w.Do(func() {

		w.authorType = authorType
	})
}

func (w *Word) GetAuthorId() kit.AuthorId {

	return w.authorId
}

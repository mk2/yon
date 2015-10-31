package word

import "github.com/mk2/yon/interp/kit"

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
	fFuncWord   = `<func> [name:%s author:%s]`
)

const (
	AuthorPrelude kit.AuthorType = "prelude"
	AuthorUser    kit.AuthorType = "user"
)

type Word struct {
	wordType   kit.WordType
	authorType kit.AuthorType
}

func (w *Word) GetWordType() kit.WordType {

	return w.wordType
}

func (w *Word) SetWordType(wordType kit.WordType) {

	w.wordType = wordType
}

func (w *Word) GetAutorType() kit.AuthorType {

	return w.authorType
}

func (w *Word) SetAuthorType(authorType kit.AuthorType) {

	w.authorType = authorType
}

package word

import (
	"sync"

	"github.com/mk2/yon/interp/kit"
)

const (
	TNilWord kit.WordType = iota
	TFuncWord
	TDictWord
	TNumberWord
	TStringWord
	TBoolWord
	TChainWord
	TArrayWord
	TNameWord
)

const (
	fNumberWord = `<number> %f`
	fStringWord = `<string> "%s"`
	fBoolWord   = `<bool>   "%t"`
	fNameWord   = `<name>   %s`
	fArrayWord  = `<array>  {%s}`
	fDictWord   = `<dict>   {%s}`
	fChainWord  = `<chain>`
	fFuncWord   = `<func>   [name:%-8s quoted:%t]`
	fNilWord    = `<nil>`
)

const (
	AuthorPrelude kit.AuthorType = "prelude"
	AuthorGo      kit.AuthorType = "go"
	AuthorUser    kit.AuthorType = "user"
)

type word struct {
	sync.Once
	wordType kit.WordType
	author   kit.Author
}

func (w *word) GetWordType() kit.WordType {

	return w.wordType
}

func (w *word) GetAuthorType() kit.AuthorType {

	return w.author.GetAuthorType()
}

func (w *word) GetAuthorId() kit.AuthorId {

	return w.author.GetAuthorId()
}

func (w *word) GetAuthor() kit.Author {

	return w.author
}

func CheckChainWord(w kit.Word) bool {

	switch w.GetWordType() {
	case TChainWord | TArrayWord | TFuncWord:
		return true
	default:
		return false
	}
}

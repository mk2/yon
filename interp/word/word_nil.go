package word

import (
	"github.com/mk2/yon/interp/author"
	"github.com/mk2/yon/interp/kit"
)

type nilWord struct {
	word
}

func NewNilWord() kit.Word {

	return &nilWord{
		word: word{wordType: TNilWord, author: author.NewPreludeAuthor()},
	}
}

func (w *nilWord) Do(m kit.Memory) (interface{}, error) {

	// do nothing

	return nil, nil
}

func (w *nilWord) String() string {

	return "nil"
}

func (w *nilWord) Format() string {

	return fNilWord
}

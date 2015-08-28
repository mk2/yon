package word
import "container/list"

type WordCode func(*Word) error

type WordType int

const (
	NotWordType WordType = iota
	EmbedWordType
	IdentWordType
	NumWordType
	StrWordType
)

type EmbedWordKind int

const (
	NotEmbedWordKind EmbedWordKind = iota
	StackOpEmbedWordKind
)

type EmbedWord struct {
	EmbedWordType EmbedWordKind
}

type IdentWord struct {
	IdentName string
}

type NumWord struct {
	Num float64
}

type StrWord struct {
	Str string
}

type Word struct {
	WordType WordType
	EmbedWord
	IdentWord
	NumWord
	StrWord
}

func (w *Word) Exec(stack *list.List) (err error) {

	switch w.WordType {

	case NumWordType:
		stack.PushFront(w.Num)

	}

	return
}
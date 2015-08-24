package yon
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
	embedWordType EmbedWordKind
}

type IdentWord struct {
	identName string
}

type NumWord struct {
	num float64
}

type StrWord struct {
	str string
}

type Word struct {
	wordType WordType
	EmbedWord
	IdentWord
	NumWord
	StrWord
}

func (w *Word) Exec(stack *list.List) error {

	switch w.wordType {

	case NumWordType:
		stack.PushFront(w.num)

	}

	return
}
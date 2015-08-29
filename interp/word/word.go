package word

type Result interface{}

type WordCode func(*BaseWord) error
type WordType int

const (
	NilWordType WordType = iota
	FuncWordType
	NumberWordType
	StringWordType
	QuoteWordType
	ArrayWordType
)

type EmbedWordKind int

const (
	NotEmbedWordKind EmbedWordKind = iota
	StackOpEmbedWordKind
)

type Word interface {
	GetWordType() WordType
	SetWordType(WordType)
}

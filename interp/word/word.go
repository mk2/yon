package word
import (
	"github.com/mk2/yon/interp/stack"
)

type Result interface{}

type WordCode func(*BaseWord) error
type WordType int

const (
	NotWordType WordType = iota
	EmbedFuncWordType
	FuncWordType
	NumberWordType
	StringWordType
)

type EmbedWordKind int

const (
	NotEmbedWordKind EmbedWordKind = iota
	StackOpEmbedWordKind
)

type Word interface {
	GetWordType() WordType
	SetWordType(WordType)
	SetStack(stack.Stack)
	GetStack() (stack.Stack, error)
	CanExec() (bool, error)
	Exec() (Result, error)
}

type EmbedWord struct {
	BaseWord
	EmbedWordType EmbedWordKind
}


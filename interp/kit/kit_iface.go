package kit

import (
	"container/list"
	"io"
)

type Word interface {
	GetWordType() WordType
	SetWordType(WordType)
	Do(m Memory) (interface{}, error)
}

type Stack interface {
	Push(v interface{}) *list.Element
	Pop() *list.Element
	Peek() *list.Element
	Print()
}

type Vocabulary interface {
	Write(k string, w Word) error
	Read(k string) Word
	LoadPrelude() error
}

type History interface {
	Leave(w Word) error
}

type Memory interface {
	Stack() Stack
	Vocab() Vocabulary
	History() History
}

type Token interface {
	GetType() TokenType
	GetPos() Position
	GetVal() string
	String() string
}

type Lexer interface {
	NextToken() Token
	GetTokens() <-chan Token
}

type Parser interface {
	NextWord() Word
	GetWords() <-chan Word
}

/*
================================================================================
IO interface
================================================================================
*/

type RuneScanner interface {
	io.RuneScanner
}

type TokenScanner interface {
	ReadToken() (Token, error)
	UnreadToken() error
}

type WordScanner interface {
	ReadWord() (Word, error)
	UnreadWord() error
}

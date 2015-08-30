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

type Memory interface {
	Stack() Stack
	Vocab() Vocabulary
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

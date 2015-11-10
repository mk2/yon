package kit

import (
	"io"
	"time"
)

// Token represents program atom
type Token interface {
	GetType() TokenType
	GetPos() Position
	GetVal() string
	String() string
}

// Lexer returns token stream
type Lexer interface {
	TokenScanner
	NextToken() Token
	GetTokens() <-chan Token
}

// Parser returns word stream
type Parser interface {
	WordScanner
	NextWord() Word
	GetWords() <-chan Word
}

// Author contains word's owner information.
type Author interface {
	GetAuthorType() AuthorType
	// GetAuthorId returns author id (mostly given random generated string)
	GetAuthorId() AuthorId
}

// Stack consists runtime temporary memory
type Stack interface {
	Push(v Word) Word
	Pop() Word
	Peek() Word
	Print()
}

// Vocabulary holds any named words
type Vocabulary interface {
	Write(k string, w Word) error
	Read(k string) Word
	LoadPrelude() error
	Print()
}

// History will contain any user operation
type History interface {
	Record(Word) error
	RecordAt(Word, time.Time) error
	Between(time.Time, time.Time) []Word
}

// Memory contains any instances of Stakc, Vocabulary, History
type Memory interface {
	Stack() Stack
	Vocab() Vocabulary
	History() History
}

// Interpreter represents abstract interpret runtimeVolabulary
type Interpreter interface {
	PrintStack()
	PrintVocab()
	PrintHistory()
	EvalAndWait(runes RuneScanner) error
	Wait() error
	Eval(runes RuneScanner) (StoppedCh, ErrorCh)
}

// Exception represents error occasion during running program
type Exception interface {
	err() error
	time() *time.Time
}

/*
================================================================================
IO interface
================================================================================
*/

// RuneScanner used for rune streaming
type RuneScanner interface {
	io.RuneScanner
}

// TokenScanner used for token streaming
type TokenScanner interface {
	ReadToken() (Token, error)
	UnreadToken() error
}

// WordScanner used for word streaming
type WordScanner interface {
	ReadWord() (Word, error)
	UnreadWord() error
}

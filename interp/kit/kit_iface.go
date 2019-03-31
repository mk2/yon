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
	Print() string
}

// Vocabulary holds any named words
type Vocabulary interface {
	// Write register the word with the class and the key.
	Write(string, string, Word) error
	// OverWrite register the word with the class and the key.
	OverWrite(string, string, Word) error
	// ReadClass returns the word searched by the given class name and key name
	ReadClass(string, string) Word
	// Read returns the word searching by the given fully qualified key
	// if the fully qualified key doesn't have any class name, search all class until the key found.
	Read(string) Word
	// LoadPrelude registers prelude words.
	LoadPrelude() error
	// NewClass makes new class. (it used for the word classification)
	NewClass(string) error
	// Print returns formatted vocabulary
	Print() string
	// Nil returns constant nil word
	Nil() Word
	// True returns constant true word
	True() Word
	// False returns constant false word
	False() Word
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
	Printf(string, ...interface{})
	Errorf(string, ...interface{})
	Println(...interface{})
	Stdout() string
	Stderr() string
}

// Interpreter represents abstract interpret runtimeVolabulary
type Interpreter interface {
	PrintStack()
	PrintVocab()
	PrintHistory()
	StdoutString() string
	StderrString() string
	EvalAndWait(runes RuneScanner) error
	SetClass(string) error
	GetClass() (string, error)
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

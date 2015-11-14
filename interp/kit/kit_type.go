package kit

// WordType is for representing word type
type WordType int
type WordFuncBody func(Memory, ...interface{}) error
type WordDoWrapper func(WordFuncBody) error

// TokenType is for representing token type
type TokenType int

// Position is unknown
type Position int

// AuthorType indicates the word author public name
type AuthorType string

// ClassType is a word cluster
type ClassType string

// AuthorId indicates the word author private name (its must be unique string)
type AuthorId string

// StoppedCh is a type for notifying the stopping.
type StoppedCh chan struct{}

// ErrorCh is a channel type for notifying the error happens.
type ErrorCh chan error

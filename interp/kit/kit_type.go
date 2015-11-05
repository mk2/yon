package kit

type WordType int
type WordFuncBody func(Memory) error

type TokenType int
type Position int

// AuthorType indicates the word author public name
type AuthorType string
// AuthorId indicates the word author private name (its must be unique string)
type AuthorId string

type StoppedCh chan struct{}
type ErrorCh chan error

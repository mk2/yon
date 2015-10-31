package kit

type WordType int
type WordFuncBody func(Memory) error

type TokenType int
type Position int

type AuthorType string

type StoppedCh chan struct{}
type ErrorCh chan error

package kit

import "container/list"

// Word represents abstract value wrapper
// words must ...
//  - Be immutable, as far as possible.
//  - Be classified-able by AuthorType
//  - Have unique author id
type Word interface {
	GetWordType() WordType
	GetAuthorType() AuthorType
	GetAuthorId() AuthorId
	GetAuthor() Author
	// String returns string representation of the word.
	String() string
	// Format returns readble formatted string of the word.
	Format() string
	Do(Memory, ...interface{}) (interface{}, error)
}

// ChainWord is basic chainable words
type ChainWord interface {
	Word
	ExtractList() *list.List
	Unshift(Word) Word
	Push(Word) Word
	Each(func(Word))
	FlatEach(func(Word))
	Size() int
}

// ArrayWord represents Array container word
type ArrayWord interface {
	ChainWord
	Put(Word)
	Array() []Word
}

type DictWord interface {
	ChainWord
	Put(Word, Word)
	Map() map[Word]Word
}

// NumberWord holds number literal
type NumberWord interface {
	Word
	Number() float64
}

// StringWord holds string literal
type StringWord interface {
	Word
}

type BoolWord interface {
	Word
	Eval() bool
}

// NameWord holds any name identifier
type NameWord interface {
	Word
	Name() string
}

// FuncWord represents any functional word (contains meta-quoted words)
type FuncWord interface {
	Word
	Name() string
}

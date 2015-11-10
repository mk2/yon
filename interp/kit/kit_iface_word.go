package kit

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
	String() string
	Do(m Memory) (interface{}, error)
}

// ChainWord is basic chainable words
type ChainWord interface {
	Word
	Push(Word) Word
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
	Tuple() map[Word]Word
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

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

type ChainWord interface {
	Word
}

// ArrayWord represents Array container word
type ArrayWord interface {
	Word
	Put(Word)
	Array() []Word
}

type NumberWord interface {
	Word
	Number() float64
}

type StringWord interface {
	Word
}

type NameWord interface {
	Word
	Name() string
}

type FuncWord interface {
	Word
	Name() string
}

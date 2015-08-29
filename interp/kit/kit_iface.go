package kit

import "container/list"

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

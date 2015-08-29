package kit

import "container/list"

type WordType int

type Word interface {
	GetWordType() WordType
	SetWordType(WordType)
	Read(m Memory) (interface{}, error)
}

type Stack interface {
	Push(v interface{}) *list.Element
	Pop() *list.Element
	Print()
}

type Vocabulary interface {
	Put(k string, w Word)
	Get(k string) Word
}

type Memory interface {
	Stack() Stack
	Vocab() Vocabulary
}

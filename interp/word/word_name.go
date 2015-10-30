package word

import "github.com/mk2/yon/interp/kit"

type NameWord struct {
	Word
	Name string
}

func NewNameWord(name string) *NameWord {

	return &NameWord{
		Word: Word{wordType: TNameWord},
		Name: name,
	}
}
func (w *NameWord) Do(m kit.Memory) (interface{}, error) {

	s := m.Stack()
	s.Push(w.Name)

	return nil, nil
}

package memory

import "github.com/mk2/yon/interp/kit"

type memory struct {
	stack   kit.Stack
	vocab   kit.Vocabulary
	history kit.History
}

func New(stack kit.Stack, vocab kit.Vocabulary, history kit.History) kit.Memory {

	return &memory{
		stack:   stack,
		vocab:   vocab,
		history: history,
	}
}

func (m *memory) Stack() kit.Stack {

	return m.stack
}

func (m *memory) Vocab() kit.Vocabulary {

	return m.vocab
}

func (m *memory) History() kit.History {

	return m.history
}

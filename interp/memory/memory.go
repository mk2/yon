package memory

import "github.com/mk2/yon/interp/kit"

type memory struct {
	kit.Memory
	stack kit.Stack
	vocab kit.Vocabulary
}

func New(stack kit.Stack, vocab kit.Vocabulary) kit.Memory {

	return &memory{
		stack: stack,
		vocab: vocab,
	}
}

func (m *memory) Stack() kit.Stack {

	return m.stack
}

func (m *memory) Vocab() kit.Vocabulary {

	return m.vocab
}

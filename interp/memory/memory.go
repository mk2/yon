package memory

import "github.com/mk2/yon/interp/kit"

type Memory struct {
	kit.Memory
	stack kit.Stack
	vocab kit.Vocabulary
}

func New(stack kit.Stack, vocab kit.Vocabulary) *Memory {

	return &Memory{
		stack: stack,
		vocab: vocab,
	}
}

func (m *Memory) Stack() kit.Stack {

	return m.stack
}

func (m *Memory) Vocab() kit.Vocabulary {

	return m.vocab
}

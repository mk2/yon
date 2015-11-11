package memory

import (
	"bytes"
	"fmt"

	"github.com/mk2/yon/interp/kit"
)

type memory struct {
	stack   kit.Stack
	vocab   kit.Vocabulary
	history kit.History
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
}

func New(stack kit.Stack, vocab kit.Vocabulary, history kit.History) kit.Memory {

	return &memory{
		stack:   stack,
		vocab:   vocab,
		history: history,
		stdout:  new(bytes.Buffer),
		stderr:  new(bytes.Buffer),
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

func (m *memory) Printf(format string, args ...interface{}) {

	m.stdout.WriteString(fmt.Sprintf(format, args...))
}

func (m *memory) Errorf(format string, args ...interface{}) {

	m.stderr.WriteString(fmt.Sprintf(format, args...))
}

func (m *memory) Println(args ...interface{}) {

	m.stdout.WriteString(fmt.Sprintln(args...))
}

func (m *memory) Stdout() string {

	s := m.stdout.String()
	m.stdout.Reset()
	return s
}

func (m *memory) Stderr() string {

	s := m.stderr.String()
	m.stderr.Reset()
	return s
}

package word

import (
	"errors"
	"github.com/mk2/yon/interp"
	"github.com/mk2/yon/interp/stack"
)

type BaseWord struct {
	Word
	wordType WordType
	stack    stack.Stack
}

func (w *BaseWord) GetWordType() WordType {

	return w.wordType
}

func (w *BaseWord) SetWordType(wordType WordType) {

	w.wordType = wordType
}

func (w *BaseWord) SetStack(stack stack.Stack) {

	w.stack = stack
}

func (w *BaseWord) GetStack() (stack.Stack, error) {

	if w.stack == nil {
		return nil, errors.New("Haven't set stack yet")
	} else {
		return w.stack, nil
	}
}

func (w *BaseWord) CanExec() (bool, error) {

	if w.stack == nil {
		return false, errors.New("the stack word is nil")
	}

	return true, nil
}

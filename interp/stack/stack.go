package stack

import "container/list"

type Stack struct {
	list.List
}

func New() *Stack {

	return &Stack{
		List: *list.New(),
	}
}

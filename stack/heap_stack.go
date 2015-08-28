package stack

import (
	"github.com/mk2/yon/stack/event"
)

type HeapStack []*event.Event

// -- sort.Interface {{{

func (s *HeapStack) Len() int {

	return len(s)
}

func (s *HeapStack) Less(i, j int) bool {

	return false
}

func (s *HeapStack) Swap(i, j int) {

}

// }}} sort.Interface

func (s *HeapStack) Push(x interface{}) {

}

func (s *HeapStack) Pop() interface{} {

	return nil
}
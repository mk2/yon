package stack

import (
	"sync"
	"github.com/mk2/yon/stack/event"
)

type Stack struct {
	sync.Mutex
	stoppingCh chan struct {}
	pushCh     chan event.Event
	stack      HeapStack

}

// New returns new stack object
func New() (s *Stack) {

	s = new(Stack)
	s.pushCh = make(chan event.Event)

	s.startPushChHandler()

	return
}

// Push is used for pushing event object on the stack
func (s *Stack) Push(event event.Event) {

	s.pushCh <- event
}

// Pop returns a object on the top of stack
func (s *Stack) Pop() event.Event {

	s.Lock()
	e := s.stack.Pop()
	s.Unlock()

	if event, ok := e.(event.Event); ok {

		return event
	}

	return nil
}

func (s *Stack) startPushChHandler() {

	go func() {

		for {
			select {

			case event := <-s.pushCh:
				s.Lock()
				s.stack.Push(event)
				s.Unlock()

			case <-s.stoppingCh:
				break

			}
		}

	}()
}

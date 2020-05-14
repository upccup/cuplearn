package stack

import "sync"

type Stack struct {
	iterms []interface{}
	lock   *sync.RWMutex
}

func New() Stack {
	s := Stack{}
	s.iterms = []interface{}{}
	s.lock = &sync.RWMutex{}
	return s
}

func (s *Stack) push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.iterms = append(s.iterms, item)
}

func (s *Stack) pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.iterms)
	if length <= 0 {
		return nil
	}

	val := s.iterms[length-1]
	s.iterms = s.iterms[0 : length-1]
	return val
}

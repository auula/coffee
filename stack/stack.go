package stack

import (
	"sync"

	"github.com/auula/coffee"
)

type Stack[V any] struct {
	size  int
	value []V
	sync.Mutex
}

func New[V any]() Stack[V] {
	return Stack[V]{
		size: 0,
	}
}

func (s *Stack[V]) Push(value V) {
	s.Lock()
	defer s.Unlock()

	s.value = append(s.value, value)
	s.size += 1
}

func (s *Stack[V]) Pop() *V {
	s.Lock()
	defer s.Unlock()

	var v V
	if s.size != 0 {
		v = s.value[s.size-1]
		s.size -= 1
		s.value = s.value[:s.size]
	}
	return &v
}

func (s *Stack[V]) Iter() coffee.Iterator[V] {
	return s
}

func (s *Stack[V]) HasNext() bool {
	s.Lock()
	defer s.Unlock()

	return s.size != 0
}

func (s *Stack[V]) Next() V {
	s.Lock()
	defer s.Unlock()

	return *s.Pop()
}

func (s *Stack[V]) Size() int {
	s.Lock()
	defer s.Unlock()

	return s.size
}

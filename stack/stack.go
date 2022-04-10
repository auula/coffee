package stack

import (
	"github.com/auula/coffee"
)

type Stack[V any] struct {
	size  int
	value []V
}

func New[V any]() Stack[V] {
	return Stack[V]{
		size: 0,
	}
}

func (s *Stack[V]) Push(value V) {
	s.value = append(s.value, value)
	s.size += 1
}

func (s *Stack[V]) Pop() *V {
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
	return s.size != 0
}

func (s *Stack[V]) Next() V {
	return *s.Pop()
}

func (s *Stack[V]) Size() int {
	return s.size
}

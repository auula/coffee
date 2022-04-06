package stack

import "github.com/auula/coffee"

type Stack[V any] struct {
	Len   uint
	value []V
}

func New[V any]() Stack[V] {
	return Stack[V]{
		Len: 0,
	}
}

func (s *Stack[V]) Push(value V) {
	s.value = append(s.value, value)
	s.Len += 1
}

func (s *Stack[V]) Pop() *V {
	var v V
	if s.Len != 0 {
		v = s.value[s.Len-1]
		s.Len -= 1
		s.value = s.value[:s.Len]
	}
	return &v
}

func (s *Stack[V]) Iter() coffee.Iterator[V] {
	return s
}

func (s *Stack[V]) HasNext() bool {
	return s.Len != 0
}

func (s *Stack[V]) Next() V {
	return *s.Pop()
}

package stack

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
	if s.Len == 0 {
		return nil
	}
	v := s.value[s.Len-1]
	if s.Len != 0 {
		s.Len -= 1
		s.value = s.value[:s.Len]
	}
	return &v
}

func (s *Stack[V]) Range(channel chan *V) {
	go func() {
		for s.Len != 0 {
			channel <- s.Pop()
		}
		close(channel)
	}()
}

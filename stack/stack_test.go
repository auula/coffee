package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New[int8]()

	for i := 0; i < 10; i++ {
		stack.Push(int8(i))
	}

	channel := make(chan *int8, 1)

	stack.Range(channel)

	for v := range channel {
		t.Log(*v)
	}
}

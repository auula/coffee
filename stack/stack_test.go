package stack

import (
	"testing"

	"github.com/auula/coffee"
)

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
func TestStackIterator(t *testing.T) {
	stack := New[int8]()

	for i := 0; i < 100; i++ {
		stack.Push(int8(i))
	}

	iter := stack.Iter()

	for iter.HasNext() {
		t.Log(iter.Next())
	}
}

func TestStackForEach(t *testing.T) {
	stack := New[int8]()

	for i := 0; i < 100; i++ {
		stack.Push(int8(i))
	}

	iter := stack.Iter()

	coffee.ForEach(iter, func(i int8) {
		t.Log(i)
	})
}

func TestStackLimit(t *testing.T) {
	stack := New[int8]()

	for i := 0; i < 100; i++ {
		stack.Push(int8(i))
	}

	iter := stack.Iter()

	coffee.Limit(iter, 3, func(i int8) {
		t.Log(i)
	})
}

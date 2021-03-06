package list

import (
	"testing"

	"github.com/auula/coffee"
)

func TestList(t *testing.T) {
	cl := New[int]()

	for i := 0; i < 10; i++ {
		cl.RPush(i)
	}

	for i := 0; i < 10; i++ {
		t.Log(cl.Get(i).Value)
	}

	cl.LPush(-1)

	t.Log(cl.Head)
	t.Log(cl.Tail)
}

func TestListInsert(t *testing.T) {
	cl := New[int]()

	for i := 0; i < 10; i++ {
		cl.RPush(i)
	}

	prevNode := cl.Get(5)
	newNode := NewNode(-5)

	newNode.Prev = prevNode
	newNode.Next = prevNode.Next
	prevNode.Next = newNode

	for i := 0; i < 12; i++ {
		t.Log(cl.Get(i))
	}
}

func TestListForEach(t *testing.T) {

	// cl := list.New[int]()
	cl := New[int]()

	for i := 0; i < 10; i++ {
		cl.RPush(i)
	}

	coffee.ForEach(cl.Iter(), func(i int) {
		t.Log(i)
	})
}

func TestListBack(t *testing.T) {
	cl := New[int]()

	for i := 0; i < 10; i++ {
		cl.RPush(i)
	}

	for i := 0; i < 10; i++ {
		t.Log(cl.Back())
	}
}

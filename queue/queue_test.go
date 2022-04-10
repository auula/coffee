package queue

import (
	"github.com/auula/coffee"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	var q Queued[int] = NewArray[int](10)

	for i := 10; i > 0; i-- {
		q.EnQueue(i)
	}

	for i := 0; i < 5; i++ {
		t.Log(q.DeQueue())
	}

	t.Log("IsFull:", q.IsFull())

	t.Log("Size:", q.Size())

	for i := 10; i < 15; i++ {
		q.EnQueue(i)
	}

	t.Log("IsFull:", q.IsFull())

	coffee.ForEach(q.Iter(), func(i int) {
		t.Log(i)
	})

	t.Log("IsFull:", q.IsFull())
}

func TestLinkedQueue(t *testing.T) {
	var q Queued[int] = New[int]()

	for i := 0; i < 10; i++ {
		q.EnQueue(i)
	}

	for i := 0; i < 10; i++ {
		t.Log(q.DeQueue())
	}

	for i := 0; i < 20; i++ {
		q.EnQueue(i)
	}

	coffee.ForEach(q.Iter(), func(i int) {
		t.Log(i)
	})
}

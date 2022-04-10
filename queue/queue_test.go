package queue

import "testing"

func TestArrayQueue(t *testing.T) {
	var q Queuer[int] = NewArray[int](10)

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

	for i := 0; i < 10; i++ {
		t.Log(q.DeQueue())
	}

	t.Log("IsFull:", q.IsFull())

}

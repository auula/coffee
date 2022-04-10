package queue

import (
	"github.com/auula/coffee"
)

type ArrayQueue[V any] struct {
	items      []V
	size       int
	head, tail int
}

func (a *ArrayQueue[V]) EnQueue(v V) bool {

	if a.tail == a.size {
		if a.head == 0 {
			return false
		}
		for i := a.head; i < a.tail; i++ {
			a.items[i-a.head] = a.items[i]
		}
		a.tail -= a.head
		a.head = 0
	}
	a.items[a.tail] = v
	a.tail += 1
	return true
}

func (a *ArrayQueue[V]) DeQueue() *V {

	if a.tail == a.head {
		return nil
	}
	v := a.items[a.head]
	a.head += 1
	return &v
}

func (a *ArrayQueue[V]) IsFull() bool {
	return a.tail == a.size && a.head == 0
}

func (a *ArrayQueue[V]) Size() int {
	return a.tail - a.head
}

func NewArray[V any](capacity int) Queued[V] {
	return &ArrayQueue[V]{
		items: make([]V, capacity),
		head:  0,
		tail:  0,
		size:  capacity,
	}
}

func (a *ArrayQueue[V]) Iter() coffee.Iterator[V] {
	return a
}

func (a *ArrayQueue[V]) HasNext() bool {
	return a.tail-a.head != 0
}

func (a *ArrayQueue[V]) Next() V {
	return *a.DeQueue()
}

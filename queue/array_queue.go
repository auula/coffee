package queue

import "sync"

type ArrayQueue[V any] struct {
	items      []V
	size       int
	head, tail int
	sync.Mutex
}

func (a *ArrayQueue[V]) EnQueue(v V) bool {
	a.Lock()
	defer a.Unlock()

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

func (a *ArrayQueue[V]) DeQueue() V {
	a.Lock()
	defer a.Unlock()

	var null V
	if a.tail == a.head {
		return null
	}
	v := a.items[a.head]
	a.head += 1
	return v
}

func (a *ArrayQueue[V]) IsFull() bool {
	a.Lock()
	defer a.Unlock()
	return a.tail == a.size && a.head == 0
}

func (a *ArrayQueue[V]) Size() int {
	a.Lock()
	defer a.Unlock()
	return a.tail - a.head
}

func NewArray[V any](capacity int) Queuer[V] {
	return &ArrayQueue[V]{
		items: make([]V, capacity),
		head:  0,
		tail:  0,
		size:  capacity,
	}
}

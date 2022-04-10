package ring

import "github.com/auula/coffee"

type RingQueue[V any] struct {
	items      []V
	head, tail int
	size       int
}

func New[V any](capacity int) RingQueue[V] {
	return RingQueue[V]{
		items: make([]V, capacity),
		size:  capacity,
	}
}

func (rq *RingQueue[V]) Put(v V) bool {
	if (rq.tail+1)%rq.size == rq.head {
		return false
	}
	rq.items[rq.tail] = v
	rq.tail = (rq.tail + 1) % rq.size
	return true
}

func (rq *RingQueue[V]) Get() *V {
	if rq.tail == rq.head {
		return nil
	}
	v := rq.items[rq.head]
	rq.head = (rq.head + 1) % rq.size
	return &v
}

func (rq *RingQueue[V]) IsFull() bool {
	return (rq.tail+1)%rq.size == rq.head
}

func (rq *RingQueue[V]) Size() int {
	return rq.tail - rq.head
}

func (rq *RingQueue[V]) Iter() coffee.Iterator[V] {
	return rq
}

func (rq *RingQueue[V]) HasNext() bool {
	return rq.tail-rq.head != 0
}

func (rq *RingQueue[V]) Next() V {
	return *rq.Get()
}

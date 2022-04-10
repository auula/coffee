package queue

import "github.com/auula/coffee"

type Queued[V any] interface {
	EnQueue(v V) bool
	DeQueue() *V
	IsFull() bool
	Size() int
	Iter() coffee.Iterator[V]
}

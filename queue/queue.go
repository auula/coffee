package queue

type Queuer[V any] interface {
	EnQueue(v V) bool
	DeQueue() V
	IsFull() bool
	Size() int
}

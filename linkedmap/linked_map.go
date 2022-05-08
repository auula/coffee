package linkedmap

type Root[T any] struct {
	Head, Tail *Node[T]
}

type Node[T any] struct {
	Prev, Next *Node[T]
	Value      T
}

// Stored structure
type Stored[T any] struct {
	LoadFactor int8
	Capacity   int16
	Hashed     func([]byte) int64
	Slot       []Root[T]
}

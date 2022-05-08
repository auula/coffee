package xmap

type Root[T any] struct {
	Head, Tail *Node[T]
}

type Node[T any] struct {
	Prev, Next *Node[T]
	Value      T
}

// Stored structure
type Stored[T any] struct {
	accessOrder bool               // the iteration ordering method
	loadFactor  int8               // map load factor
	capacity    int16              // current value size
	hashed      func([]byte) int64 // hash function
	slot        []Root[T]          // data slot
}

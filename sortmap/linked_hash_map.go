package sortmap

import "github.com/auula/coffee"

type Node[K comparable, V any] struct {
	Prev, Next *Node[K, V]
	Key        K
	Value      V
}

// Stored structure
type LinkedHashMap[K comparable, V any] struct {
	capacity   int               // total capacity
	head, tail *Node[K, V]       // linked list
	table      map[K]*Node[K, V] // data storeage
	size       int               // current size
}

func New[K comparable, V any](capacity int) LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		capacity: capacity,
		table:    make(map[K]*Node[K, V], capacity),
	}
}

func (hashmap *LinkedHashMap[K, V]) Put(key K, value V) bool {

	node := &Node[K, V]{
		Key:   key,
		Value: value,
	}

	if hashmap.size == 0 {
		hashmap.head = node
		hashmap.tail = node
		hashmap.table[key] = node
		hashmap.size += 1
		return true
	}

	if node, ok := hashmap.table[key]; ok {
		node.Value = value
		return true
	}

	addNodeAtTail(hashmap, node)
	hashmap.table[key] = node
	hashmap.size += 1

	return true
}

func (hashmap *LinkedHashMap[K, V]) Remove(key K) {
	if node, ok := hashmap.table[key]; ok {
		moveNode(hashmap, node)
		hashmap.size -= 1
		delete(hashmap.table, key)
	}
}

func (hashmap *LinkedHashMap[K, V]) Get(key K) V {

	var (
		node *Node[K, V]
		ok   bool
	)

	if node, ok = hashmap.table[key]; !ok {
		return node.Value
	}

	moveNode(hashmap, node)
	addNodeAtTail(hashmap, node)

	return node.Value
}

func (hashmap *LinkedHashMap[K, V]) Clear() {
	hashmap.size = 0
	hashmap.head = nil
	hashmap.tail = nil
	hashmap.table = make(map[K]*Node[K, V], hashmap.capacity)
}

func (hashmap *LinkedHashMap[K, V]) Size() int {
	return hashmap.size
}

// 从两个节点中间删除节点
func moveNode[K comparable, V any](hashmap *LinkedHashMap[K, V], node *Node[K, V]) {
	if node == hashmap.head {
		node.Next.Prev = nil
		hashmap.head = node.Next
		return
	}
	if node == hashmap.tail {
		node.Prev.Next = nil
		hashmap.tail = node.Prev
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// addTail 添加节点到链表尾巴
func addNodeAtTail[K comparable, V any](hashmap *LinkedHashMap[K, V], node *Node[K, V]) {
	node.Prev = hashmap.tail
	hashmap.tail.Next = node
	hashmap.tail = node
}

func (hashmap *LinkedHashMap[K, V]) Iter() coffee.Iterator[V] {
	return hashmap
}

func (hashmap *LinkedHashMap[K, V]) HasNext() bool {
	return hashmap.size != 0
}

func (hashmap *LinkedHashMap[K, V]) Next() V {
	var node *Node[K, V]
	if node = hashmap.head; node != nil {
		hashmap.head = hashmap.head.Next
		hashmap.size -= 1
		return node.Value
	}
	return node.Value
}

func (hashmap *LinkedHashMap[K, V]) Range(fn func(Node[K, V])) {
	var node *Node[K, V]
	for hashmap.size != 0 {
		if node = hashmap.head; node != nil {
			hashmap.head = hashmap.head.Next
			hashmap.size -= 1
			fn(*node)
		}
	}
}

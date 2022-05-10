package coffee

type Node[V any] struct {
	Prev, Next *Node[V]
	Value      V
}

// Stored structure
type LinkedHashMap[K comparable, V any] struct {
	accessOrder bool           // the iteration ordering method
	capacity    int            // total capacity
	head, tail  *Node[V]       // linked list
	table       map[K]*Node[V] // data storeage
	size        int            // current size
}

func NewLinkedHashMap[K comparable, V any](capacity int, accessOrder bool) LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		accessOrder: accessOrder,
		capacity:    capacity,
		table:       make(map[K]*Node[V], capacity),
	}
}

func (hashmap *LinkedHashMap[K, V]) Put(key K, value V) bool {

	node := &Node[V]{
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
		moveNode(node)
		node.Next = nil
		addNodeAtTail(hashmap, node)
		return true
	}

	addNodeAtTail(hashmap, node)
	hashmap.table[key] = node
	hashmap.size += 1

	return true
}

func (hashmap *LinkedHashMap[K, V]) Remove(key K) {
	if node, ok := hashmap.table[key]; ok {
		moveNode(node)
		delete(hashmap.table, key)
	}
}

func (hashmap *LinkedHashMap[K, V]) Get(key K) *V {

	var (
		node *Node[V]
		ok   bool
	)

	if node, ok = hashmap.table[key]; !ok {
		return nil
	}

	moveNode(node)
	addNodeAtTail(hashmap, node)

	return &node.Value
}

func (hashmap *LinkedHashMap[K, V]) Clear() {
	hashmap.size = 0
	hashmap.head = nil
	hashmap.tail = nil
	hashmap.table = make(map[K]*Node[V], hashmap.capacity)
}

func (hashmap *LinkedHashMap[K, V]) Size() int {
	return hashmap.size
}

// 从两个节点中间删除节点
func moveNode[V any](node *Node[V]) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// addTail 添加节点到链表尾巴
func addNodeAtTail[K comparable, V any](hashmap *LinkedHashMap[K, V], node *Node[V]) {
	node.Prev = hashmap.tail
	hashmap.tail.Next = node
	hashmap.tail = node
}

func (hashmap *LinkedHashMap[K, V]) Iter() Iterator[V] {
	return hashmap
}

func (hashmap *LinkedHashMap[K, V]) HasNext() bool {
	return hashmap.size != 0
}

func (hashmap *LinkedHashMap[K, V]) Next() V {
	var node *Node[V]
	if node = hashmap.head; node != nil {
		hashmap.head = hashmap.head.Next
		hashmap.size -= 1
		return node.Value
	}
	return node.Value
}

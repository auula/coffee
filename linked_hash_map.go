package coffee

const (
	// offset64 FNVa offset basis.
	// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value.
	// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

// Sum64 gets the string and returns its uint64 hash value.
// a new 64-bit FNV-1a Hashed which makes no memory allocations.
// Its Sum64 method will lay the value out in big-endian byte order.
// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function
func Sum64(key []byte) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}

type Node[T any] struct {
	Prev, Next *Node[T]
	Value      T
}

// Stored structure
type LinkedHashMap[T any] struct {
	accessOrder bool                // the iteration ordering method
	capacity    int                 // total capacity
	head, tail  *Node[T]            // linked list
	hashed      func([]byte) uint64 // hash function
	table       map[uint64]*Node[T] // data storeage
	size        int                 // current size
}

func NewLinkedHashMap[T any](capacity int, accessOrder bool) LinkedHashMap[T] {
	return LinkedHashMap[T]{
		accessOrder: accessOrder,
		capacity:    capacity,
		hashed:      Sum64,
		table:       make(map[uint64]*Node[T], capacity),
	}
}

func (hashmap *LinkedHashMap[T]) Put(key string, value T) bool {

	sum64 := Sum64([]byte(key))

	node := &Node[T]{
		Value: value,
	}

	if hashmap.size == 0 {
		hashmap.head = node
		hashmap.tail = node
		return true
	}

	if node, ok := hashmap.table[sum64]; ok {
		node.Value = value
		moveNode(node)
		addNodeAtTail(hashmap, node)
		hashmap.size += 1
		return true
	}

	addNodeAtTail(hashmap, node)
	hashmap.table[sum64] = node

	return true
}

func (hashmap *LinkedHashMap[T]) Remove(key T) {

}

func (hashmap *LinkedHashMap[T]) Get(key T) *T {
	return nil
}

func (hashmap *LinkedHashMap[T]) Cap() int {
	return 0
}

// 从两个节点中间删除节点
func moveNode[T any](node *Node[T]) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// addTail 添加节点到链表尾巴
func addNodeAtTail[T any](hashmap *LinkedHashMap[T], node *Node[T]) {
	node.Prev = hashmap.tail
	node.Next = nil
	hashmap.tail.Next = node
	hashmap.tail = node
}

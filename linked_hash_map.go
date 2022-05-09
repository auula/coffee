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

type Node[K, V any] struct {
	Prev, Next *Node[T]
	Key        K
	Value      T
}

// Stored structure
type LinkedHashMap[K, V any] struct {
	accessOrder bool                   // the iteration ordering method
	capacity    int                    // total capacity
	head, tail  *Node[K, V]            // linked list
	hashed      func([]byte) uint64    // hash function
	table       map[uint64]*Node[K, V] // data storeage
	size        int                    // current size
}

func NewLinkedHashMap[K, V any](capacity int, accessOrder bool) LinkedHashMap[T] {
	return LinkedHashMap[K, V]{
		accessOrder: accessOrder,
		capacity:    capacity,
		hashed:      Sum64,
		table:       make(map[uint64]*Node[K, V], capacity),
	}
}

func (hashmap *LinkedHashMap[K, V]) Put(key K, value V) bool {

	sum64 := Sum64([]byte(key))

	node := &Node[K, V]{
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
		node.Next = nil
		addNodeAtTail(hashmap, node)
		hashmap.size += 1
		return true
	}

	addNodeAtTail(hashmap, node)
	hashmap.table[sum64] = node
	hashmap.size += 1

	return true
}

func (hashmap *LinkedHashMap[K, V]) Remove(key K) {

}

func (hashmap *LinkedHashMap[K, V]) Get(key K) *V {
	return nil
}

func (hashmap *LinkedHashMap[K, V]) Size() int {
	return hashmap.size
}

// 从两个节点中间删除节点
func moveNode[K, V any](node *Node[K, V]) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// addTail 添加节点到链表尾巴
func addNodeAtTail[K, V any](hashmap *LinkedHashMap[K, V], node *Node[K, V]) {
	node.Prev = hashmap.tail
	hashmap.tail.Next = node
	hashmap.tail = node
}

package heap

func Sort[T any](tree []T, fn func(x, y T) bool) {
	Build(tree, fn, len(tree))
	for i := len(tree) - 1; i >= 0; i-- {
		tree[0], tree[i] = tree[i], tree[0]
		Heapify(tree, fn, i, 0)
	}
}

func Build[T any](tree []T, fn func(x, y T) bool, len int) {
	parent := (len - 1) / 2
	for parent >= 0 {
		Heapify(tree, fn, len, parent)
		parent -= 1
	}
}

func Heapify[T any](tree []T, fn func(x, y T) bool, len, i int) {

	if i >= len {
		return
	}

	var (
		max         = i
		left, right = i*2 + 1, i*2 + 2
	)
	if left < len && fn(tree[left], tree[max]) {
		max = left
	}
	if right < len && fn(tree[right], tree[max]) {
		max = right
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		Heapify(tree, fn, len, max)
	}
}

type Tree[T any] struct {
	item []T
	fn   func(x, y T) bool
}

func New[T any](fn func(x, y T) bool, v ...T) Tree[T] {
	var tree Tree[T]
	tree.fn = fn
	tree.item = v
	Build(tree.item, tree.fn, len(tree.item))
	return tree
}

func (h *Tree[T]) Put(e T) {
	h.item = append(h.item, e)
	Build(h.item, h.fn, len(h.item))
}

func (h *Tree[T]) Peek() *T {
	return &h.item[0]
}

func (h *Tree[T]) Poll() *T {
	element := &h.item[0]
	h.item = h.item[1:]
	Build(h.item, h.fn, len(h.item))
	return element
}

func (h *Tree[T]) Size() int {
	return len(h.item)
}

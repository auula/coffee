package heap

import "github.com/auula/coffee"

func Sort[T coffee.Number](tree []T, fn func(x, y T) bool) {
	Build(tree, fn, len(tree))
	for i := len(tree) - 1; i >= 0; i-- {
		tree[0], tree[i] = tree[i], tree[0]
		Heapify(tree, fn, i, 0)
	}
}

func Build[T coffee.Number](tree []T, fn func(x, y T) bool, len int) {
	parent := (len - 1) / 2
	for parent >= 0 {
		Heapify(tree, fn, len, parent)
		parent -= 1
	}
}

func Heapify[T coffee.Number](tree []T, fn func(x, y T) bool, len, i int) {

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

func New[T coffee.Number](arr []T) {

}

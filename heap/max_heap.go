package heap

import "github.com/auula/coffee"

func Sort[T coffee.Number](tree []T) {
	Build(tree, len(tree))
	for i := len(tree) - 1; i >= 0; i-- {
		tree[0], tree[i] = tree[i], tree[0]
		Heapify(tree, i, 0)
	}
}

func Build[T coffee.Number](tree []T, len int) {
	parent := (len - 1) / 2
	for parent >= 0 {
		Heapify(tree, len, parent)
		parent -= 1
	}
}

func Heapify[T coffee.Number](tree []T, len int, i int) {

	if i >= len {
		return
	}

	var (
		max         = i
		left, right = i*2 + 1, i*2 + 2
	)
	if left < len && tree[left] > tree[max] {
		max = left
	}
	if right < len && tree[right] > tree[max] {
		max = right
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		Heapify(tree, len, max)
	}
}

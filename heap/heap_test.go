package heap

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "max",
			args: args{
				tree: []int{5, 4, 1, 3, 2, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.tree, func(x, y int) bool {
				return x > y
			})
			t.Log(tt.args.tree)
		})
	}
}

func TestHeap(t *testing.T) {
	heap := New(func(x, y int) bool {
		return x > y
	}, 5, 4, 1, 3, 2, 10)
	heap.Put(100)
	t.Log(*heap.Peek()) // 100 √
	heap.Poll()
	t.Log(*heap.Peek()) // 10 √
}

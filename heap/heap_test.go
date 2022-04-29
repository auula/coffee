package heap

import "testing"

func TestSort(t *testing.T) {
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

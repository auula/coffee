package sort

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		sequence []int
		val      int
		len      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successful",
			args: args{
				sequence: []int{10, 11, 34, 45, 77, 78, 98, 122},
				val:      45,
				len:      8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.sequence, tt.args.val, tt.args.len); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

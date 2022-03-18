package sort

import "testing"

func TestSequenceByInt64(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int64
	}{
		{
			name:     "int64",
			sequence: []int64{12123, 22, 121, 2, 34, 5665, 1213},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sequence(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByFloat64(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{121.23, 2.2, 12.1, 2, 34, 5.665, 121.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sequence(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

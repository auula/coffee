package sort

import "testing"

func TestSequenceByCounting(t *testing.T) {
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
			Counting(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

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
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sequence(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceBySorting(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sorting(func(i, j float64) bool {
				return j > i
			}, tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByQuick(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByBubble(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bubble(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByMerge(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByShell(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Shell(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceBySelection(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Selection(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

func TestSequenceByInsertion(t *testing.T) {
	tests := []struct {
		name     string
		sequence []float64
	}{
		{
			name:     "float64",
			sequence: []float64{3.3, 2.2, 1.1, 2, 3, 5.665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insertion(tt.sequence)
			t.Log(tt.sequence)
		})
	}
}

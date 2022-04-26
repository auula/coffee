package sort

import (
	"github.com/auula/coffee"
)

func Sequence[N coffee.Number](sequence []N) {
	Merge(sequence)
}

func Sorting[N coffee.Number](cp func(i, j N) bool, sequence []N) {
	for gap := len(sequence) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(sequence); i++ {
			previous := i - gap
			current := sequence[i]
			for previous >= 0 && cp(current, sequence[previous]) {
				sequence[previous+gap] = sequence[previous]
				previous -= gap
			}
			if previous+gap != i {
				sequence[previous+gap] = current
			}
		}
	}
}

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
			pervIndex := i - gap
			current := sequence[i]
			for pervIndex >= 0 && cp(current, sequence[pervIndex]) {
				sequence[pervIndex+gap] = sequence[pervIndex]
				pervIndex -= gap
			}
			if pervIndex+gap != i {
				sequence[pervIndex+gap] = current
			}
		}
	}
}

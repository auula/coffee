package sort

import (
	"github.com/auula/coffee"
)

// Shell sorting function
func Shell[N coffee.Number](sequence []N) {
	var (
		previous int
		current  N
	)
	for gap := len(sequence) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(sequence); i++ {
			previous = i - gap
			current = sequence[i]
			for previous >= 0 && current < sequence[previous] {
				sequence[previous+gap] = sequence[previous]
				previous -= gap
			}
			if previous+gap != i {
				sequence[previous+gap] = current
			}
		}
	}
}

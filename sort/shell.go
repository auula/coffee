package sort

import (
	"github.com/auula/coffee"
)

// Shell sorting function
func Shell[N coffee.Number](sequence []N) {
	var (
		pervIndex int
		current   N
	)
	for gap := len(sequence) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(sequence); i++ {
			pervIndex = i - gap
			current = sequence[i]
			for pervIndex >= 0 && current < sequence[pervIndex] {
				sequence[pervIndex+gap] = sequence[pervIndex]
				pervIndex -= gap
			}
			if pervIndex+gap != i {
				sequence[pervIndex+gap] = current
			}
		}
	}
}

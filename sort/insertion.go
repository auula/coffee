package sort

import (
	"github.com/auula/coffee"
)

// Insertion sorting function
func Insertion[N coffee.Number](sequence []N) {
	for i := 1; i < len(sequence); i++ {
		pervIndex := i - 1
		current := sequence[i]
		for pervIndex >= 0 {
			if current < sequence[pervIndex] {
				sequence[pervIndex+1] = sequence[pervIndex]
			} else {
				break
			}
			pervIndex -= 1
		}
		sequence[pervIndex+1] = current
	}
}

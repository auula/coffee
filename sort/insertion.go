package sort

import (
	"github.com/auula/coffee"
)

// Insertion sorting function
func Insertion[N coffee.Number](sequence []N) {
	for i := 1; i < len(sequence); i++ {
		pervIndex := i - 1
		current := sequence[i]
		// 每次检测左边排好的并且插入
		for pervIndex >= 0 && current < sequence[pervIndex] {
			sequence[pervIndex+1] = sequence[pervIndex]
			pervIndex--
		}
		if pervIndex+1 != i {
			sequence[pervIndex+1] = current
		}
	}
}

package sort

import (
	"github.com/auula/coffee"
)

// Selection sorting function
func Selection[N coffee.Number](sequence []N) {
	for i := 0; i < len(sequence)-1; i++ {
		// find max index
		max := i
		for j := i + 1; j < len(sequence); j++ {
			if sequence[max] > sequence[j] {
				max = j
			}
		}
		// exchange
		sequence[max], sequence[i] = sequence[i], sequence[max]
	}
}

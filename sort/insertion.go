package sort

import (
	"github.com/auula/coffee"
)

// Insertion sorting function
func Insertion[N coffee.Number](sequence []N) {
	for i := 1; i < len(sequence); i++ {
		previous := i - 1
		current := sequence[i]
		for previous >= 0 {
			if current < sequence[previous] {
				sequence[previous+1] = sequence[previous]
			} else {
				break
			}
			previous -= 1
		}
		sequence[previous+1] = current
	}
}

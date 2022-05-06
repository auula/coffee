package sort

import (
	"github.com/auula/coffee"
)

func Search[T coffee.Number](sequence []T, target T, len int) int {
	low, high, middle := 0, len-1, 0
	for low < high {
		middle = low + ((high - low) >> 1)
		if sequence[middle] == target {
			return middle
		}
		if sequence[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

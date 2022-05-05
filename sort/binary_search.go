package sort

import (
	"github.com/auula/coffee"
)

func Search[T coffee.Number](sequence []T, val T, len int) int {
	low, high, middle := 0, len-1, 0
	for low <= high {
		middle = low + (high-low)>>1
		if sequence[middle] == val {
			return middle
		}
		if sequence[middle] < val {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

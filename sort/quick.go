package sort

import "github.com/auula/coffee"

func Quick[N coffee.Number](sequence []N) {
	quickSort(sequence, 0, len(sequence)-1)
}

func partition[N coffee.Number](sequence []N, startIndex, endIndex int) int {
	left, right := startIndex, endIndex
	pivot := sequence[startIndex]
	for left != right {
		for left < right && sequence[right] > pivot {
			right -= 1
		}
		for left < right && sequence[left] <= pivot {
			left += 1
		}
		if left < right {
			sequence[left], sequence[right] = sequence[right], sequence[left]
		}
	}
	sequence[startIndex], sequence[left] = sequence[left], pivot
	return left
}

func quickSort[N coffee.Number](sequence []N, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(sequence, startIndex, endIndex)
	quickSort(sequence, startIndex, pivotIndex-1)
	quickSort(sequence, pivotIndex+1, endIndex)
}

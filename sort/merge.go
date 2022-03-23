package sort

import "github.com/auula/coffee"

// Merge sorting function
func Merge[N coffee.Number](sequence []N) {
	copy(sequence, mergeSort(sequence))
}

// group arrays
func mergeSort[N coffee.Number](sequence []N) []N {
	length := len(sequence)
	if length < 2 {
		return sequence
	}
	middle := length / 2
	return sort(mergeSort(sequence[:middle]), mergeSort(sequence[middle:]))
}

// group arrays sorting
func sort[N coffee.Number](left, right []N) []N {
	var result []N
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}

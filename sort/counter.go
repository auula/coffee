package sort

import "github.com/auula/coffee"

func Counter[N coffee.Integer](sequence []N) {
	copy(sequence, counter(sequence))
}

func counter[N coffee.Integer](sequence []N) []N {
	max, min := sequence[0], sequence[0]

	for i := 1; i < len(sequence); i++ {
		if sequence[i] > max {
			max = sequence[i]
		}
		if sequence[i] < min {
			min = sequence[i]
		}
	}

	difference := max - min + 1
	countArray := make([]int, difference)
	for i := 0; i < len(sequence); i++ {
		countArray[sequence[i]-min] += 1
	}

	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	sortedArray := make([]N, len(sequence))
	for i := len(sequence) - 1; i >= 0; i-- {
		sortedArray[countArray[sequence[i]-min]-1] = sequence[i]
		countArray[sequence[i]-min] -= 1
	}

	return sortedArray
}

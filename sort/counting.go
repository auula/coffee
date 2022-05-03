package sort

import "github.com/auula/coffee"

func Counting[N coffee.Integer](sequence []N) {
	copy(sequence, counter(sequence))
}

func counter[N coffee.Integer](sequence []N) []N {
	max, min := sequence[0], sequence[0]

	for i := 1; i < len(sequence); i++ {
		if sequence[i] > max {
			max = N(i)
		}
		if sequence[i] < min {
			min = N(i)
		}
	}

	offset := max - min + 1

	count := make([]N, offset)

	for i := 0; i < len(sequence); i++ {
		count[sequence[i]-min]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	sorted := make([]N, len(sequence))

	for i := len(sequence) - 1; i >= 0; i-- {
		sorted[count[sequence[i]-min]-1] = sequence[i]
		count[sequence[i]-min]--
	}

	return sorted
}

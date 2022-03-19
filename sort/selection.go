package sort

// Selection sorting function
func Selection[N Number](sequence []N) {
	for i := 0; i < len(sequence)-1; i++ {
		max := i
		for j := i + 1; j < len(sequence); j++ {
			if sequence[max] > sequence[j] {
				max = j
			}
		}
		sequence[max], sequence[i] = sequence[i], sequence[max]
	}
}

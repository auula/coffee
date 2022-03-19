package sort

// Insertion sorting function
func Insertion[N Number](sequence []N) {
	for i := 1; i < len(sequence)-1; i++ {
		pervIndex := i - 1
		current := sequence[i]
		for pervIndex >= 0 && sequence[pervIndex] > current {
			sequence[pervIndex+1] = sequence[pervIndex]
			pervIndex--
		}
		if pervIndex+1 != i {
			sequence[pervIndex+1] = current
		}
	}
}

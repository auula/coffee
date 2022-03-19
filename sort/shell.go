package sort

// Shell sorting function
func Shell[N Number](sequence []N) {
	var (
		pervIndex int
		current   N
	)
	for gap := len(sequence) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(sequence)-1; i++ {
			pervIndex = i - gap
			current = sequence[i]
			for pervIndex >= 0 && sequence[pervIndex] > current {
				sequence[pervIndex+gap] = sequence[pervIndex]
				pervIndex -= gap
			}
			if pervIndex+gap != i {
				sequence[pervIndex+gap] = current
			}
		}
	}
}

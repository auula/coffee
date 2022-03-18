package sort

func Bubble[N Number](sequence []N) {
	for i := 0; i < len(sequence)-1; i++ {
		for j := 0; j < len(sequence)-i-1; j++ {
			if sequence[j] > sequence[j+1] {
				sequence[j], sequence[j+1] = sequence[j+1], sequence[j]
			}
		}
	}
}

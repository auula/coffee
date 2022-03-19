package sort

type Number interface {
	comparable
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

type Sorting[N Number] interface {
	Sequence([]N)
}

func Sequence[N Number](sequence []N) {
	Selection(sequence)
}

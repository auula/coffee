package sort

import (
	"github.com/auula/coffee"
)

type Sorting[N coffee.Number] interface {
	Sequence([]N)
}

func Sequence[N coffee.Number](sequence []N) {
	Shell(sequence)
}

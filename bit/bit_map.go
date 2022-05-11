package bit

type Map struct {
	bytes []int16
	nbits int
}

func New(nbits int) Map {
	return Map{
		nbits: nbits,
		bytes: make([]int16, nbits/16+1),
	}
}

func (m *Map) Set(index int) {
	if index > m.nbits {
		return
	}
	byteIndex := index / 16
	bitIndex := index % 16
	m.bytes[byteIndex] |= (1 << bitIndex)
}

func (m *Map) Peek(index int) bool {
	if index > m.nbits {
		return false
	}
	byteIndex := index / 16
	bitIndex := index % 16
	return (m.bytes[byteIndex] & (1 << bitIndex)) != 0
}

package coffee

type Number interface {
	comparable
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

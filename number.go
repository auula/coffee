package coffee

type Integer interface {
	comparable
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

type Float interface {
	comparable
	float32 | float64
}

type Number interface {
	comparable
	Integer | Float
}

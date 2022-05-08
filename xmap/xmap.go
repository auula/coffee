package xmap

// ModeType is storage mode
type ModeType int8

const (
	Array  ModeType = 1 // Array storage mode
	Linked              // LinkedList storage mode
)

// Mapping structure
type Mapping struct {
	LoadFactor int8
	Capacity   int16
	Hashed     func([]byte) int64
	Mode       int
}

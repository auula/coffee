package xmap

// Mode is storage mode
type Mode int8

const (
	Array  Mode = 1 // Array storage mode
	Linked          // LinkedList storage mode
)

// Mapping structure
type Map struct {
	LoadFactor int8
	Capacity   int16
	Hashed     func([]byte) int64
	Store      Mode
}

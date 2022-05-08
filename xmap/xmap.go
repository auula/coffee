package xmap

// Mode is storage mode
type Mode int8

const (
	Array  Mode = 1 // Array storage mode
	Linked          // LinkedList storage mode
)

// Stored structure
type Stored struct {
	LoadFactor int8
	Capacity   int16
	Hashed     func([]byte) int64
	Data       Mode
}

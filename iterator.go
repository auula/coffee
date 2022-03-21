package coffee

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func ForEach[T any](iterator Iterator[T], fn func(T)) {
	var element T
	for iterator.HasNext() {
		element = iterator.Next()
		fn(element)
	}
}

func Map[T, V any](iterator Iterator[T], fn func(T) V) []V {
	var (
		element T
		result  []V
	)
	for iterator.HasNext() {
		element = iterator.Next()
		result = append(result, fn(element))
	}
	return result
}

func Filter[T any](iterator Iterator[T], fn func(T) bool) []T {
	var (
		element T
		vector  []T
	)
	for iterator.HasNext() {
		element = iterator.Next()
		if !fn(element) {
			vector = append(vector, element)
		}
	}
	return vector
}

func Reduce[T, R any](iterator Iterator[T], start R, fn func(T, R) R) R {
	result := start
	for iterator.HasNext() {
		element := iterator.Next()
		result = fn(element, result)
	}
	return result
}

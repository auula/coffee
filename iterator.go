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

func Filter[T any](iterator Iterator[T], flag bool, fn func(T) bool) []T {
	var (
		element T
		vector  []T
	)
	for iterator.HasNext() {
		element = iterator.Next()
		choose := fn(element)
		if (choose && flag) || (!choose && !flag) {
			vector = append(vector, element)
		}
	}
	return vector
}

func Reduce[T, R any](iterator Iterator[T], start R, fn func(R, T) R) R {
	result := start
	for iterator.HasNext() {
		element := iterator.Next()
		result = fn(result, element)
	}
	return result
}

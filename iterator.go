package coffee

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func ForEach[T any](iterator Iterator[T], fn func(T)) {
	for iterator.HasNext() {
		fn(iterator.Next())
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
		slice   []T
	)
	for iterator.HasNext() {
		element = iterator.Next()
		choose := fn(element)
		if (choose && flag) || (!choose && !flag) {
			slice = append(slice, element)
		}
	}
	return slice
}

func Reduce[T, R any](iterator Iterator[T], start R, fn func(R, T) R) R {
	result := start
	for iterator.HasNext() {
		element := iterator.Next()
		result = fn(result, element)
	}
	return result
}

func Limit[T any](iterator Iterator[T], n int, fn func(T)) {
	var count int
	for iterator.HasNext() {
		if count == n {
			break
		}
		fn(iterator.Next())
		count++
	}
}

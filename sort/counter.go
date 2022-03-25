package sort

import "github.com/auula/coffee"

func Counter[N coffee.Number](sequence []N) {
	copy(sequence, counter(sequence))
}

// 拿到数列里面最大的值，初始化一个计数用的数组
func counter[N coffee.Number](sequence []N) []N {
	return []N{}
}

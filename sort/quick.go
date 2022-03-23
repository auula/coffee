package sort

import "github.com/auula/coffee"

func Quick[N coffee.Number](sequence []N) {
	quickSort(sequence, 0, len(sequence)-1)
}

// 把一组无序数列的通过基准进行分割
// 然后通过基准值排序
// 小的放左边，大的放右边，然后递归操作
func partition[N coffee.Number](sequence []N, startIndex, endIndex int) int {
	left, right := startIndex, endIndex
	pivot := sequence[startIndex]
	for left != right {
		// 如果右边的元素大于基准元素就向左移动指针
		for left < right && sequence[right] > pivot {
			right -= 1
		}
		// 如果左边的元素小于基准元素就向左移动指针
		for left < right && sequence[left] <= pivot {
			left += 1
		}
		// 如果上面两个指针移动都停止了就交换指针上的元素
		if left < right {
			sequence[left], sequence[right] = sequence[right], sequence[left]
		}
	}
	// 如果两个指针重合就交换基准指针所在位置的元素
	sequence[startIndex], sequence[left] = sequence[left], pivot
	return left
}

// 把一段数据通过基准值不断切分并且排序
func quickSort[N coffee.Number](sequence []N, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(sequence, startIndex, endIndex)
	quickSort(sequence, startIndex, pivotIndex-1)
	quickSort(sequence, pivotIndex+1, endIndex)
}

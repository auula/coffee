//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/auula/coffee"
	"github.com/auula/coffee/stack"
)

func main() {
	s := stack.New[int]()

	for i := 1; i <= 100; i++ {
		s.Push(i)
	}

	sum := 0

	coffee.ForEach(s.Iter(), func(i int) {
		sum += i
	})

	fmt.Println(s.Pop())
	fmt.Println(sum)
}

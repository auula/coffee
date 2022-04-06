//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/auula/coffee"
	"github.com/auula/coffee/list"
)

func main() {

	ls := list.New[int]()

	for i := 0; i < 10; i++ {
		ls.RPush(i)
	}

	fmt.Println("Front:", ls.Front())

	for i := 0; i < 20; i++ {
		ls.LPush(i)
	}

	fmt.Println("Back:", ls.Front())

	coffee.ForEach(ls.Iter(), func(i int) {
		fmt.Println(i)
	})

}

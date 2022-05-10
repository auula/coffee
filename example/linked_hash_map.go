package main

import (
	"fmt"

	"github.com/auula/coffee"
)

func main() {

	var lmap = coffee.NewLinkedHashMap[string, float64](10, true)

	lmap.Put("Go", 91.2)
	lmap.Put("Java", 100)
	lmap.Put("Rust", 80.1)

	coffee.ForEach(lmap.Iter(), func(i float64) {
		fmt.Println(i)
	})
}

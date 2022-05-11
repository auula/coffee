package main

import (
	"fmt"

	"github.com/auula/coffee/sortmap"
)

func main() {

	var sm = sortmap.New[string, float64](10)

	sm.Put("Go", 91.2)
	sm.Put("Java", 100)
	sm.Put("Rust", 80.1)

	// coffee.ForEach(sm.Iter(), func(i float64) {
	// 	fmt.Println(i)
	// })

	sm.Range(func(node sortmap.Node[string, float64]) {
		fmt.Printf("\n key:%s \t value:%f \n", node.Key, node.Value)
	})
}

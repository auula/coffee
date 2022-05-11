package main

import (
	"fmt"

	"github.com/auula/coffee"
)

func main() {

	var m = coffee.NewLinkedHashMap[string, float64](10)

	m.Put("Go", 91.2)
	m.Put("Java", 100)
	m.Put("Rust", 80.1)

	m.Put("Go", 92.2)

	// coffee.ForEach(lmap.Iter(), func(i float64) {
	// 	fmt.Println(i)
	// })

	m.Range(func(node coffee.Node[string, float64]) {
		fmt.Printf("\n key:%s \t value:%f \n", node.Key, node.Value)
	})
}

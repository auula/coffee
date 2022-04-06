# coffee
General programming toolkit for Golang.

---

### Quick Start

[1. List](#List)

[2. Stack](#Stack)

---

## List

LinkedList enhanced operation API:

```go
package main

import (
	"fmt"

	"github.com/auula/coffee"
	"github.com/auula/coffee/list"
)

func main() {

	ls := list.New[int]()

	for i := 0; i < 10; i++ {
		ls.RPush(i) // Insert from the right
	}

	fmt.Println("Front:", ls.Front()) // Get element from head

	for i := 0; i < 20; i++ {
		ls.LPush(i) // Insert from the left
	}

	fmt.Println("Back:", ls.Back()) // Get element from tail

	coffee.ForEach(ls.Iter(), func(i int) {
		fmt.Println(i)
	})

}
```
## Stack

Stack enhanced operation API:

```go
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
```
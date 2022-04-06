# coffee
General programming toolkit for Golang.

---

### Quick Start

[1. List](#List)

[2. Stack](#Stack)

---

### List

Linked List basic api:

```go
func main() {
    cl := list.New[int]()

	for i := 0; i < 10; i++ {
		cl.RPush(i) // Insert from the right
	}

    for i := 0; i < 10; i++ {
		cl.LPush(i) // Insert from the left
	}

    // Get element from head
    t.Log(cl.Front())

    // Get element from tail
    t.Log(cl.Back())

    // Get node by index
    node := cl.Get(6) 
    t.Log(node)
}
```
Linked List enhanced operation api:

```go
    channel := make(chan int, 1)

	cl.Range(channel)

	for v := range channel {
		t.Log(v)
	}
```
OR
```go
    coffee.ForEach(cl.Iter(), func(i int) {
		t.Log(i)
	})
```


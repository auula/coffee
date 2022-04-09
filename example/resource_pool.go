//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/auula/coffee/pool"
)

type connection struct {
	id int32
}

func (c connection) Close() error {
	fmt.Println("closed connection,id = ", c.id)
	return nil
}

var (
	count   int32
	wg      = new(sync.WaitGroup)
	factory = func() (connection, error) {
		atomic.AddInt32(&count, 1)
		return connection{id: count}, nil
	}
)

func main() {
	p := pool.New(factory, 5)

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			connect, _ := p.Acquire()
			fmt.Println(connect.id)
			time.Sleep(time.Duration(rand.Intn(1000) * int(time.Microsecond)))
			p.Release(connect)
		}(i)
	}

	wg.Wait()
	p.Close()
}

package pool

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type connection struct {
	id int32
}

func (c connection) Close() error {
	fmt.Println("closed connection,id = ", c.id)
	return nil
}

var count int32

func TestPool(t *testing.T) {
	p := New(
		func() (*connection, error) {
			atomic.AddInt32(&count, 1)
			return &connection{id: count}, nil
		},
		5,
	)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			connect, _ := p.Acquire()
			t.Log(i, ":", connect.id)
			p.Release(*connect)
			time.Sleep(time.Duration(rand.Intn(1000) * int(time.Microsecond)))
		}(i)
	}

	wg.Wait()
	p.Close()
}

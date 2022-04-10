package pool

import (
	"errors"
	"io"
	"sync"

	"github.com/auula/coffee"
)

type Resource interface {
	io.Closer
}

var (
	ErrPoolClosed      = errors.New("pool has been closed")
	ErrPoolAcquireFail = errors.New("pool acquire fail")
)

type Pool[T Resource] struct {
	lock      sync.Mutex
	resources chan T
	factory   func() (T, error)
	closed    bool
}

func New[T Resource](factory func() (T, error), size int) Pool[T] {
	if size <= 0 {
		size = 10
	}
	return Pool[T]{
		factory:   factory,
		resources: make(chan T, size),
	}
}

func (p *Pool[T]) Acquire() (T, error) {
	var null T
	select {
	case resource, ok := <-p.resources:
		if !ok {
			return null, ErrPoolClosed
		}
		return resource, nil
	default:
		return p.factory()
	}
}

func (p *Pool[T]) Release(resource T) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		resource.Close()
	}

	select {
	case p.resources <- resource:
	default:
		resource.Close()
	}
}

func (p *Pool[T]) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for resource := range p.resources {
		resource.Close()
	}

}

func (p *Pool[T]) Iter() coffee.Iterator[T] {
	return p
}

func (p *Pool[T]) HasNext() bool {
	return !p.closed
}

func (p *Pool[T]) Next() T {
	var null T
	if res, err := p.Acquire(); err == nil {
		return res
	}
	return null
}

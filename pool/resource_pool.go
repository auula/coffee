package pool

import (
	"errors"
	"io"
	"sync"
)

type Resource interface {
	io.Closer
}

var ErrPoolClosed = errors.New("Pool has been closed.")

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
	select {
	case resource, ok := <-p.resources:
		if !ok {
			var null T
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

	close(p.resources)

	for resource := range p.resources {
		resource.Close()
	}

}

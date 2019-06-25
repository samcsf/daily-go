// GO in action example - pool pattern
package pool

import (
	"errors"
	"io"
	"math/rand"
	"sync"
	"time"
)

type FactoryFn func(n int) io.Closer
type Pool struct {
	locker    sync.Mutex
	closed    bool
	resources chan io.Closer
	factory   FactoryFn
}

// create resources
// save resource when close the resouce
func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(poolSize int, factory FactoryFn) (*Pool, error) {
	if poolSize < 0 {
		return nil, errors.New("Size is too small")
	}
	return &Pool{
		resources: make(chan io.Closer, poolSize),
		factory:   factory,
	}, nil
}

func (p *Pool) Acquire() io.Closer {
	select {
	case r := <-p.resources:
		return r
	default:
		return p.factory(rand.Int())
	}
}

func (p *Pool) Release(r io.Closer) {
	p.locker.Lock()
	defer p.locker.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		return
	default:
		r.Close()
	}
}

func (p *Pool) Close() {
	p.locker.Lock()
	defer p.locker.Unlock()

	p.closed = true

	// 不close掉的话，for range 一个通道会等到天荒地老
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}

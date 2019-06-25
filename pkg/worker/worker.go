package worker

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	worker chan Worker
	wg     sync.WaitGroup
}

func New(poolSize int) *Pool {
	p := &Pool{
		worker: make(chan Worker, poolSize),
	}

	p.wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		go func() {
			for w := range p.worker {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *Pool) Shutdown() {
	close(p.worker)
	p.wg.Wait()
}

func (p *Pool) Run(works ...Worker) {
	for _, work := range works {
		p.worker <- work
	}
}

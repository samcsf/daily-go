// GO实战中的并发例子runner
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type TaskFunc func(int) error

type Runner struct {
	duration    time.Duration
	timeout     chan time.Time
	osInterrupt chan os.Signal
	done        chan error
	tasks       []TaskFunc
}

func New(duration time.Duration) *Runner {
	return &Runner{
		duration:    duration,
		timeout:     make(chan time.Time),
		osInterrupt: make(chan os.Signal, 1),
		done:        make(chan error),
	}
}

func (r *Runner) AddTask(fns ...TaskFunc) {
	r.tasks = append(r.tasks, fns...)
}

func (r *Runner) Start() error {
	// subscribe system signal to r.osInterrupt
	signal.Notify(r.osInterrupt, os.Interrupt)

	// submit goroutine to handle timeout
	go func() {
		time := <-time.After(r.duration)
		r.timeout <- time
	}()

	// submit the tasks
	go func() {
		for id, task := range r.tasks {
			if r.gotInterupt() {
				r.done <- errors.New("Got OS interrupt!")
				return
			}
			task(id)
		}
		r.done <- nil
	}()

	select {
	case <-r.timeout:
		return errors.New("Timeout!!")
	case err := <-r.done:
		return err
	}
}

func (r *Runner) gotInterupt() bool {
	select {
	case <-r.osInterrupt:
		signal.Stop(r.osInterrupt)
		return true
	default:
		return false
	}
}

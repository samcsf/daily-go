package worker

import (
	"fmt"
	"testing"
	"time"
)

type MyTask struct {
	id int
}

func (t *MyTask) Task() {
	n := 20 * time.Millisecond
	fmt.Printf("Task %d will take %d ms\n", t.id, n/time.Millisecond)
	time.Sleep(n)
	fmt.Printf("Task %d finished\n", t.id)
}

// common case
func TestCommon(t *testing.T) {
	pool := New(2)

	go func() {
		n := 50 * time.Millisecond
		time.Sleep(n)
		t.Errorf("Task not finished in time")
	}()

	pool.Run(&MyTask{1}, &MyTask{2}, &MyTask{3})
	pool.Shutdown()
}

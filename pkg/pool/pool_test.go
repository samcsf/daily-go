package pool

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type mockDBConn struct {
	id int
}

func NewDBConn(id int) io.Closer {
	fmt.Printf("Id: %d conn allocated\n", id)
	return &mockDBConn{id}
}

func (c *mockDBConn) Close() error {
	fmt.Printf("Id: %d conn destroied\n", c.id)
	return nil
}

// common case
func TestCommon(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pool, _ := New(2, NewDBConn)

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			<-time.After(time.Duration(rand.Intn(5)) * time.Second)
			conn := pool.Acquire()
			c, _ := conn.(*mockDBConn)
			fmt.Printf("Id: %d conn used\n", c.id)
			pool.Release(conn)
			wg.Done()
		}()
	}
	wg.Wait()
	pool.Close()
}

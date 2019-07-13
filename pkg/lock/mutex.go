package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// mutex has 2 mode, normal & starve
	// when in normal mode new coming goroutine will competed with waiting goroutine
	// but if a goroutine wait more than 1ms, it will change to starve mode
	// when in starve mode the waiter will get the lock in order
	var m sync.Mutex
	go lockItWithLongTime(m)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go grabTheLock(i, m, wg)
	}
	wg.Wait()
}

func lockItWithLongTime(m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	<-time.After(100 * time.Millisecond)
}

func lockItWithShortTime(m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	<-time.After(8e5 * time.Nanosecond)
}

func grabTheLock(id int, m sync.Mutex, wg sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()
	fmt.Printf("Id: %d grab the lock\n", id)
	wg.Done()
}

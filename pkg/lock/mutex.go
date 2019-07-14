// 知识层面不够，暂时放弃，想复现饥饿模式的变现
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
	m := &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(10)
	start := time.Now()
	// just make sure all the compete codes are ready
	go lockItWithLongTime(m)
	for i := 0; i < 10; i++ {
		go grabTheLock(i, m, &wg)
	}
	m.Lock()
	// <-time.After(500000 * time.Nanosecond)
	<-time.After(2 * time.Millisecond)
	fmt.Println("???")
	m.Unlock()
	wg.Wait()
	fmt.Println(time.Now().Sub(start).String())
}

func lockItWithLongTime(m *sync.Mutex) {
	(*m).Lock()
	defer (*m).Unlock()
	<-time.After(1000 * time.Millisecond)
}

func lockItWithShortTime(m *sync.Mutex) {
	(*m).Lock()
	defer (*m).Unlock()
	<-time.After(2e5 * time.Nanosecond)
}

func grabTheLock(id int, m *sync.Mutex, wg *sync.WaitGroup) {
	(*m).Lock()
	defer (*m).Unlock()
	fmt.Printf("Id: %d grab the lock\n", id)
	(*wg).Done()
}

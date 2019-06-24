package runner

import (
	"fmt"
	"strings"
	"syscall"
	"testing"
	"time"
)

func task(id int) error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Task ", id, " running")
	return nil
}

func task2(id int) error {
	time.Sleep(1 * time.Second)
	fmt.Println("Task ", id, " running")
	return nil
}

// timeout will cause timeout error
func TestTimeout(t *testing.T) {
	r := New(20 * time.Millisecond)
	r.AddTask(task, task, task)
	err := r.Start()
	if err != nil {
		errStr := err.Error()
		expect := "Timeout!!"
		if strings.Compare(errStr, expect) != 0 {
			t.Errorf("Expect '%s', Got '%s'", expect, errStr)
		}
	} else {
		t.Errorf("Time out should return an timeout error")
	}
}

// common case
func TestCommon(t *testing.T) {
	r := New(40 * time.Millisecond)
	r.AddTask(task, task, task)
	err := r.Start()
	if err != nil {
		t.Errorf("Should not got error, Got: \n %v", err)
	}
}

// system interupt
func TestInterrupt(t *testing.T) {
	r := New(200 * time.Second)
	r.AddTask(task2, task2, task2)

	go func() {
		time.After(1 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	err := r.Start()
	if err != nil {
		errStr := err.Error()
		expect := "Got OS interrupt!"
		if strings.Compare(errStr, expect) != 0 {
			t.Errorf("Expect '%s', Got '%s'", expect, errStr)
		}
	} else {
		t.Errorf("Time out should return an timeout error")
	}
}

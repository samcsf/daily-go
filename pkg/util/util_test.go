package util

import (
	"runtime"
	"testing"
)

func TestPrintMem(t *testing.T) {
	PrintMem()

	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums = append(nums, i)
	}

	PrintMem()

	nums = nil
	runtime.GC()

	PrintMem()
}

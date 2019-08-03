package util

import (
	"log"
	"runtime"
)

// Gin style shortcut
type H map[string]interface{}

func ChkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}
}

/*
	打印内存信息
	MemStats
		Alloc         从heap中分配的内存, = HeapAlloc
		TotalAlloc    累计从heap中分配的内存, 与Alloc不同的是free后不会减掉
		Lookups       runtime累计指针查找次数
		Mallocs       累计allocate的heap对象数
		Frees         累计释放的heap对象数，存活heap对象数 = Mallocs - Frees

		HeapSys       heap bytes from os, heapSys = heapIdle + heapInuse
		PauseTotalNs  累计GC时间(每次GC会停止所有goroutine执行，这个参数可以留意GC对执行的影响)
		PauseNs[256]  一个256的数组，循环记录每次GC时间，PauseNs[(NumGC+255)%256]取最后GC时间
		PauseEnd[256] 一个256的数组，循环记录每次GC结束时间
		NumGC         GC次数
		NumForcedGC   强制GC次数
*/
func PrintMemDetail() {
	stat := runtime.MemStats{}
	runtime.ReadMemStats(&stat)
	log.Printf("Alloc: %5.2f\n", float64(stat.Alloc>>10))
	log.Printf("TotalAlloc: %5.2f\n", float64(stat.TotalAlloc>>10))
	log.Printf("Mallocs: %5.2f\n", float64(stat.Mallocs))
	log.Printf("Frees: %5.2f\n", float64(stat.Alloc))
	log.Printf("HeapIdle: %5.2f\n", float64(stat.HeapIdle>>10))
	log.Printf("HeapInuse: %5.2f\n", float64(stat.HeapInuse>>10))
	log.Printf("NumGC: %5.2f\n", float64(stat.NumGC))
	log.Println("-------------Sys------------------")
	log.Printf("Sys: %5.2f\n", float64(stat.Sys>>10))
	log.Printf("HeapSys: %5.2f\n", float64(stat.HeapSys>>10))
	log.Printf("StackSys: %5.2f\n", float64(stat.StackSys>>10))
	log.Printf("MSpanSys: %5.2f\n", float64(stat.MSpanSys>>10))
	log.Printf("MCacheSys: %5.2f\n", float64(stat.MCacheSys>>10))
	log.Printf("BuckHashSys: %5.2f\n", float64(stat.BuckHashSys>>10))
	log.Printf("GCSys: %5.2f\n", float64(stat.GCSys>>10))
	log.Printf("OtherSys: %5.2f\n", float64(stat.OtherSys>>10))
	log.Println("----------------------------------")
}

func PrintMem() {
	stat := runtime.MemStats{}
	runtime.ReadMemStats(&stat)
	log.Printf("Alloc: %5.2fk\t TotalAlloc: %5.2fk\t NumGC: %d \n", float64(stat.Alloc>>10), float64(stat.TotalAlloc>>10), stat.NumGC)
}

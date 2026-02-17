package main

import (
	"fmt"
	"runtime"
)

type HugeStruct struct {
	data [1024]byte // 1KB の構造体
}

func printMemStats(label string) {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%-28s] Alloc: %8d KB  |  Sys: %8d KB  |  HeapInuse: %8d KB  |  HeapReleased: %8d KB\n",
		label, m.Alloc/1024, m.Sys/1024, m.HeapInuse/1024, m.HeapReleased/1024)
}

func main() {
	printMemStats("初期状態")

	m := make(map[int]*HugeStruct)
	for i := 0; i < 1_000_000; i++ {
		m[i] = &HugeStruct{}
	}
	printMemStats("100万件挿入後")

	for k := range m {
		delete(m, k)
	}
	printMemStats("全件 delete 後")
	fmt.Printf("len(m) = %d\n", len(m))

	m = nil
	printMemStats("m = nil 後")
}

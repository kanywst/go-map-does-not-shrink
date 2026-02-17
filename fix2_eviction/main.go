package main

import (
	"fmt"
	"runtime"
)

type HugeStruct struct {
	data [1024]byte
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%-36s] Alloc: %8d KB  |  HeapInuse: %8d KB  |  Sys: %8d KB\n",
		label, m.Alloc/1024, m.HeapInuse/1024, m.Sys/1024)
}

func main() {
	fmt.Println("=== Fix 2: Eviction (Copy to New Map) ===")
	printMemStats("Initial State")

	m := make(map[int]HugeStruct)
	for i := 0; i < 1_000_000; i++ {
		m[i] = HugeStruct{}
	}
	printMemStats("After Adding 1M Items")

	// 100件だけ残してdelete
	for k := range m {
		if k >= 100 {
			delete(m, k)
		}
	}
	runtime.GC()
	printMemStats("After Reducing to 100 items + GC")

	// 回避策②: 新Mapへコピー（Eviction）
	newMap := make(map[int]HugeStruct, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	m = nil
	m = newMap
	runtime.GC()
	printMemStats("After Eviction (Copy) + GC")
	_ = m
}

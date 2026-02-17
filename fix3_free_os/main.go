package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type HugeStruct struct {
	data [1024]byte
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%-36s] Alloc: %8d KB  |  HeapInuse: %8d KB  |  HeapReleased: %8d KB  |  Sys: %8d KB\n",
		label, m.Alloc/1024, m.HeapInuse/1024, m.HeapReleased/1024, m.Sys/1024)
}

func main() {
	fmt.Println("=== Fix 3: debug.FreeOSMemory() ===")
	printMemStats("Initial State")

	m := make(map[int]HugeStruct)
	for i := 0; i < 1_000_000; i++ {
		m[i] = HugeStruct{}
	}
	printMemStats("After Adding 1M Items")

	m = nil
	printMemStats("After m = nil (Before GC)")

	runtime.GC()
	printMemStats("After GC (Before FreeOSMemory)")

	debug.FreeOSMemory()
	printMemStats("After debug.FreeOSMemory()")
}

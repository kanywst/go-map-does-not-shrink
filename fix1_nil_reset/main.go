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
	fmt.Println("=== Fix 1: m = nil + make (Reset) ===")
	printMemStats("Initial State")

	m := make(map[int]HugeStruct)
	for i := 0; i < 1_000_000; i++ {
		m[i] = HugeStruct{}
	}
	printMemStats("After Adding 1M Items")

	// 回避策①: nil リセット + 再作成
	m = nil
	runtime.GC()
	printMemStats("After m = nil + GC")

	m = make(map[int]HugeStruct, 100)
	for i := 0; i < 100; i++ {
		m[i] = HugeStruct{}
	}
	printMemStats("After Re-make + 100 items")
	_ = m
}

package main

/*
#include <mach/mach.h>
#include <mach/task_info.h>

unsigned long long get_rss() {
    struct mach_task_basic_info info;
    mach_msg_type_number_t count = MACH_TASK_BASIC_INFO_COUNT;
    kern_return_t kr = task_info(mach_task_self(), MACH_TASK_BASIC_INFO,
                                (task_info_t)&info, &count);
    if (kr != KERN_SUCCESS) return 0;
    return info.resident_size;
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type HugeStruct struct {
	data [1024]byte
}

func getRSS() uint64 {
	return uint64(C.get_rss())
}

func printAll(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	rss := getRSS()
	fmt.Printf("[%-36s] Alloc: %8d KB  |  Sys: %8d KB  |  HeapReleased: %8d KB  |  RSS: %8d KB\n",
		label, m.Alloc/1024, m.Sys/1024, m.HeapReleased/1024, rss/1024)
}

func main() {
	printAll("初期状態")

	m := make(map[int]HugeStruct)
	for i := 0; i < 1_000_000; i++ {
		m[i] = HugeStruct{}
	}
	printAll("100万件挿入後")

	for k := range m {
		delete(m, k)
	}
	printAll("全件 delete 後 (GC前)")

	runtime.GC()
	printAll("GC 後")

	m = nil
	runtime.GC()
	printAll("m = nil + GC 後")

	debug.FreeOSMemory()
	printAll("debug.FreeOSMemory() 後")
}

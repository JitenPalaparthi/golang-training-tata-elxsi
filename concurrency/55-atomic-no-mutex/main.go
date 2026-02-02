package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Do not communicate by sharing memory; instead, share memory by communicating.

var counter int64 = 0

func main() {

	wg := new(sync.WaitGroup)

	//defer wg.Wait()

	//wg.Add(1)
	//go func() {

	for i := 1; i < 10000; i++ {
		wg.Add(2)
		go Incr(wg)
		go Decr(wg)
	}

	wg.Wait()
	fmt.Println("Counter:", atomic.LoadInt64(&counter)) // This would ensure that loading the value with a proper memory order
}

func Incr(wg *sync.WaitGroup) {
	atomic.AddInt64(&counter, 1)
	wg.Done()
}

func Decr(wg *sync.WaitGroup) {
	atomic.AddInt64(&counter, -1)
	wg.Done()
}

// global data that is shared by multiple threads/ multiple goroutines
// atomic operations are by design thread safe
// they are direct single instruction operations on the processor
// but they have limited functionality
// can be used for global counters or add sub etc..

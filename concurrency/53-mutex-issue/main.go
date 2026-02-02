package main

import (
	"fmt"
	"sync"
)

// Do not communicate by sharing memory; instead, share memory by communicating.

var counter = 0

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
	//	wg.Done()
	//}()

	// wg.Add(1)
	// go func() {
	// 	for i := 1; i < 1000; i++ {
	// 		wg.Add(1)
	// 		go Decr(wg)
	// 	}
	// 	wg.Done()
	// }()

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func Incr(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}

func Decr(wg *sync.WaitGroup) {
	counter--
	wg.Done()
}

// global data that is shared by multiple threads/ multiple goroutines

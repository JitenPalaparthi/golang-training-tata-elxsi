package main

import (
	"math/rand/v2"
	"time"
)

func main() { // p3

	go func() {
		println("Hello World")
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			go DoWork()
		}
	}()

	println("Hello Golang minds!")

	time.Sleep(time.Millisecond * 1) // deliberately blocking the main gorountines
	// This is not a solution. It may run , may not all gogoroutines

}

func DoWork() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += rand.IntN(31231)
	}
	println("Sum:", sum)
}

// Main is also a goroutine
// That means main would be poised to run on a P(Processor but not OS processor, it is a wrapper on top of a thread.. Thread + additional stuff)
// To run a goroutine just use a keyword go infront of a call
// The Main goroutine does not wait for other goroutines to complete its executuon

// Each P has a run queue capasity of 256
// the local queue is full, it would be placed on Global Queue
// When the P finishes its work , The load balancer would send the goroutines from Global queue to Local queue of a P
//

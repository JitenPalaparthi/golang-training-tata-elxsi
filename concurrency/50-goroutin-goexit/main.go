package main

import (
	"math/rand/v2"
	"runtime"
	"time"
)

func main() { // p3

	go func() { // goroutine1
		println("Hello World")
	}()

	go func() { // goroutine2
		i := 1
		for {
			if i >= 100 {
				runtime.Goexit() // This would not crash, bcz it is not main ..
			}
			go DoWork()
			i++
		}
	}()

	println("Hello Golang minds!")
	time.Sleep(time.Second * 2)

	//os.Exit(0)
	runtime.Goexit() // Is this a solution? not exactly
	// before going to exit the caller goroutine, it checks are there any other goroutines runnig spawn in the caller
	// before going to it exit the caller, it waits for all other goroutines to complete their exexcution (all others means w.r.t caller)
	// on main it crashes after executing all other goroutines
	// if called on non main , it would exit that goroutine
	println("os exit is called")
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
// There is not concept of parent and child goroutines.. may thing a caller is a goroutine and for the callee it could be a parent but not technically
// goroutines do not guarantee the order of execution
// Each P has a run queue capasity of 256
// the local queue is full, it would be placed on Global Queue
// When the P finishes its work , The load balancer would send the goroutines from Global queue to Local queue of a P
//

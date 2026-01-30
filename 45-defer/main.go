package main

import (
	"fmt"
)

func main() {
	var fibs []int
	println("Hello another main")

	defer fmt.Println("End of main") // deferred to be executed to the end of the caller

	defer fmt.Println(fibs)

	defer func(fibs []int) {
		fibs = fibo(10)
	}(fibs)

	func() { // func1
		defer func() { // func1.1
			println("Hello Tata Elixe Minds!")
		}()
		println("start of func1")
	}() // main.main.func1.func1.1

	fmt.Println("Start of main")

}

func fibo(num uint) []int {
	a, b := 0, 1
	var fib []int
	for i := 0; i < int(num); i++ {
		fib = append(fib, a)
		a, b = b, a+b
	}
	return fib
}

// defer defers the execution of the function/method
// when there multiple derfers for a caller , all deferred calles are stored in a stack(FILO)
// when they execute they pops out
// every caller would maintain its own stack
// deferred functionsa re 100% executed even during panics

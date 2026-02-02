package main

import (
	"demo/fileops" // init called here
	"fmt"
	"log"
	"math/rand/v2"
	"runtime"

	"sync"
)

var (
	wg       *sync.WaitGroup
	fileName string = "data.txt"
)

/// init is a special functions like main. It is calld before even main is executed..
// can create any number of init functions
// can create init in any package
// cannot give the order of execution of init
// can think it as package level constructor

func init() {
	wg = new(sync.WaitGroup)
	fileops.FileName = fileName
}

func init() {
	println("just for demo purpose--1")
}

func init() {
	println("just for demo purpose--2")
}

func init() {
	println("just for demo purpose--2")
}

func main() {
	defer wg.Wait()
	//wg.Add(3) // can give directly number of goroutines
	wg.Add(1) //+1
	go func() {
		defer wg.Done()
		if lines, err := fileops.ReadFileIntoSlice(); err != nil {
			log.Println(err)
			runtime.Goexit()
		} else {
			fmt.Println(lines)
		}

	}()
	wg.Add(1) // +1
	go func() {
		defer wg.Done()
		fibs := fib(10)
		fmt.Println(fibs)
	}()
	wg.Add(1) //+1
	go func() {
		primes := Primenubmers(10, 20)
		fmt.Println(primes)
		wg.Done() // -1
	}()
	wg.Add(1)
	go func() { // goroutine2
		i := 1
		for {
			if i >= 100 {
				wg.Done()
				runtime.Goexit() // This would not crash, bcz it is not main ..
			}
			wg.Add(1)
			go func() {
				DoWork() // how can I put done here?
				wg.Done()
			}()
			//go DoWorkWg(wg)
			i++
		}
	}()

	// //runtime.Goexit()
	// time.Sleep(time.Second * 2)

	//wg.Wait() // wait until the waitgroup state is 0

}

func fib(r int) []int {
	a, b := 0, 1
	fibs := make([]int, 0)
	for i := 1; i <= r; i++ {
		fibs = append(fibs, a)
		a, b = b, a+b
	}
	return fibs
}

func Primenubmers(start, end int) []int {
	primes := []int{}
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func DoWork() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += rand.IntN(31231)
	}
	println("Sum:", sum)
}

func DoWorkWg(wg *sync.WaitGroup) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += rand.IntN(31231)
	}
	println("Sum:", sum)
	wg.Done()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	. "os" // using . can call functions and types directlyh without using package name
	"runtime"
	"time"
)

func main() {

	go func() {

		if file, err := Open("data.txt"); err != nil { // os.Open
			log.Println(err.Error())
			runtime.Goexit()

		} else {
			defer file.Close()
			lines := make([]string, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() { // as long as the scanner.Scan gives true
				line := scanner.Text()
				lines = append(lines, line)
			}
			fmt.Println(lines)
		}

	}()

	go func() {
		fibs := fib(10)
		fmt.Println(fibs)
	}()

	go func() {
		primes := Primenubmers(10, 20)
		fmt.Println(primes)
	}()

	//runtime.Goexit()
	time.Sleep(time.Second * 2)

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

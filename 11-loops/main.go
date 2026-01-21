package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	println()

	println("Like a while loop")

	num := 1

	for num <= 10 { // just a condition like a while loop

		if num%2 == 0 {
			print(num, " ")
		}

		num++
	}
	println()
	println("Using continue")
	num = 1

	for ; num <= 10; num++ {

		if num%2 == 0 {
			continue
		}
		print(num, " ")
	}

	println()
	println("Unconditional loop")

	num = 1

	for { // the loop is unconditional, still can use break
		if num > 10 {
			break
		}
		if num%3 == 0 {
			print(num, " ")
		}
		num++
	}

	println()

	println("Nested loops")
	count := 0
	for i := 1; i <= 5; i++ {
		for j := 1; ; j++ {
			if j == 3 {
				break // The basic break breask only the loop that has been called
			}
			println("i:", i, "j:", j)
			count++
		}
	}
	println("Counter:", count)

	println("Nested loops , break both inner and outer loops")
	count = 0
outer:
	for i := 1; i <= 5; i++ {
		for j := 1; ; j++ {
			if j == 3 {
				break outer // break with label
			}
			println("i:", i, "j:", j)
			count++
		}
	}
	println("Counter:", count)

	println("break outer using switch case")
exit:
	for {
		num := rand.IntN(100000)
		switch {
		case num%2 == 0:
			println(num, "is divisible by 2")
			fallthrough
		case num%17 == 0:
			if num%17 == 0 {
				println(num, "is divisible by 17 and breaking the loop")
				break exit
			}
		}
	}

	println("Loop with multiple varialbles")

	for i, j := 1, 10; i <= 5; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	a, b := 0, 1

	println("fibonacci numbers")

	for {
		print(a, " ")
		a, b = b, a+b
		if a >= 50 {
			break
		}
	}

}

// for loop
// for range loop

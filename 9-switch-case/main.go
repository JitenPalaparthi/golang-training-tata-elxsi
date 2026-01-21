package main

import "math/rand/v2"

func main() {

	day := uint8(4)

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Satturday")
	default:
		println("Noday")
	}

	// empty switch

	num := rand.IntN(90000)

	switch { // empty switch
	case num%2 == 0:
		println(num, "is even number")
	default:
		println(num, "is odd number")
	}

	switch { // empty switch
	case num >= 0 && num <= 100:
		println(num, "is 0-100")
	case num > 100 && num <= 200:
		println(num, "is 101-200")
	case num > 200:
		println(num, "is greater than 200")
	default:
		println(num, "is negative number")
	}

	char := 'k'

	switch char { // switch with multiple cases in a single case
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is vovel")
	default:
		println(string(char), "either consonent or a special char")
	}

	num = 44

	switch {
	case num%8 == 0:
		{
			println(num, "is divisible by 8")
		}
		fallthrough
	case num%4 == 0:
		{
			println(num, "is divisible by 4")
		}
		fallthrough
	case num%2 == 0:
		{
			println(num, "is divisible by 2")
		}
	}

	println("false negative using fallthrough")

	switch {
	case num%2 == 0:
		{
			println(num, "is divisible by 2")
		}
		fallthrough
	case num%4 == 0:
		{
			println(num, "is divisible by 4")
		}
		fallthrough
	case num%8 == 0:
		{
			println(num, "is divisible by 8")
		}
	}

}

package main

func main() {
	num := 1
loop:
	if num%2 == 0 {
		println(num)
	} else {
		goto odd
	}
	num++
	if num < 10 {
		goto loop
	}

odd:
	if num <= 10 {
		println(num)
		num++
		goto loop
	}
}

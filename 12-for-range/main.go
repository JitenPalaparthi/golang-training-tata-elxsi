package main

import "strconv"

func main() {

	for i := range 10 { // comparatively new one , not there before 1.22
		println(i)
	}

	str := "Hello World" // Collection of bytes

	for i := 0; i < len(str); i++ {
		print(string(str[i]), " <> ")
	}

	println()

	str = "Hello 世界" // It is collection of bytes.. Always when you len. it is the lenof bytes
	for i := 0; i < len(str); i++ {
		print(string(str[i]))
	}
	println()

	for i, v := range str { // When range loop is used on string, instead of iterating byte by byte , it iterates char by char
		println(i, "-->", string(v))
	}

	for _, v := range str { // When range loop is used on string, instead of iterating byte by byte , it iterates char by char
		println("-->", string(v))
	}

	for i := range str { //for i,_ := range str
		println(i)
	}

	str = "934734837123934"

	// Sum of each digit
	sum := 0
	for _, v := range str {
		if num, err := strconv.Atoi(string(v)); err == nil {
			sum += num
		} else {
			println(err.Error())
			sum = 0
			break
		}
	}

	println("Sum of str:", sum)

}

// Hello ä¸ç

// range loop is used on collections, maps, channels
// for i := range 10 --> it gives the value of i
// on array, string, slice --> It gives index and value
// On map --> It gives key and value
// On Channel --> It gives the value received from the channel

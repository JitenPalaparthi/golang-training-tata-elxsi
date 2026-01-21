package main

import (
	"strconv"
)

func main() {
	age := uint8(19)
	if age >= 18 { // the ultimate is true or false
		println("eligible for vote")
	} else {
		println("not eligible for vote")
	}

	str := "123f1223"

	if num, err := strconv.Atoi(str); err != nil {
		println(err.Error())
	} else {
		println(num)
	}

	//println(num)

	age, gender := 18, 'F' // int32 or rune

	if age >= 18 && (gender == 'F' || gender == 'f') {
		println("she is eligible for marriage, according to India")
	} else if age >= 21 && (gender == 'M' || gender == 109) {
		println("he is eligible for marriage, accodring to India")
	} else {
		println("invalid gender")
	}

	//char := rand.IntN(256) // It gives random numbers less than 256

}

// Arthimetic
// Compare
// Logical --> && ||

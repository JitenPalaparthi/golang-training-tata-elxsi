package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint8 = 254

	var num2 uint16 = uint16(num1)

	println(num2)

	var num3 uint = 323233232232386764 // 1000111110001011010111010011000001011010000 00010000 11001100
	//																					 10000 11001100
	//									  1000111110001011010111010011000001011010000 00010000 11001100
	var num5 int = -323233232232386764

	var num4 uint16 = uint16(num3)
	var num6 uint16 = uint16(num5)

	fmt.Printf("%b\n", num3)
	fmt.Printf("%b\n", num5)
	fmt.Printf("%d %b\n", num4, num4)
	fmt.Printf("%d %b\n", num6, num6)

	var f1 float32 = 23232.23

	var num7 int = int(f1)
	println(num7)

	// ok1 := true

	// var num8 uint8 = uint8(ok1) // cannot cast it , only numbders can be casted among

	// rune

	var char1 rune // int32 --> 4bytes --> utf8

	char1 = 'a'
	char1 = 79
	char1 = '世'

	var num10 int32 = 132322

	num10 = num10 + char1

	// byte --> uint8

	// type integer = int // type alias

	// var i integer = 1312232

	// var char2 int32 = 'A'
	// var char3 uint8 = 'B'

	println(string(char1)) //19990

	char2 := 12312

	println(string(char2))

	str1 := string(char2) // type cast a number is casted as a char
	println(str1)

	str2 := strconv.Itoa(char2) // type conversion
	fmt.Println(str2, "Type of str2:", reflect.TypeOf(str2))

	str3 := fmt.Sprint(char2) // Formatter, converts everything or anything to string

	str4 := fmt.Sprint("Hello", true, 1231, false, "world", str3) // Format function or a method
	println(str4)

	a, s, m, d := calc(20, 10)
	println(a, s, m, d)

	a1, _, m1, d1 := calc(20, 10) // black identifier
	println(a1, m1, d1)

	_, b1, _, _ := calc(20, 10)
	println(b1)

	str5 := "123a12312"

	i1, err := strconv.Atoi(str5) // what is err ? err is an interface, Error()string
	if err != nil {               // There is some error
		println(err.Error())
	} else {
		println(i1)
	}

	str6 := "9833sfs433"
	i2, _ := strconv.Atoi(str6)
	println(i2)

	str7 := "tr13213ue"
	bo1, _ := strconv.ParseBool(str7)
	println(bo1)

}

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

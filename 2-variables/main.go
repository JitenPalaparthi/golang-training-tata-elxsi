package main

import (
	"fmt"
	"reflect"
)

func main() { // Stack Frame

	var num1 int       // Zero Value 0
	var num2 uint8     // Zero Value 0
	var float1 float32 // Zero Value 0
	var ok1 bool       // Zero Value false
	var str1 string    // Zero Value ""

	var num3 int = 12312 // No Zero value , bcz value has been assigned to the variable

	num1 = 54343
	num2 = 123
	float1 = 4564.454
	ok1 = true
	str1 = "Hello World!"

	fmt.Println(num1, num2, num3, float1, ok1, str1)

	var num4 = 12312             // Type inference
	var float2 = 97878.566423423 // Type inference
	var ok2 = true               // Type inference
	var str2 = "Hello Tata"      // Type inference
	var age uint8 = 41           // Dont want age to be int(8 bytes), it should be 1 byte so not used type inference
	fmt.Println(num4, float2, ok2, str2, age)

	fmt.Printf("%d %.3f %v %s 0X%X\n", num4, float2, ok2, str2, age)
	fmt.Printf("%d %.3f %v %s %b\n", num4, float2, ok2, str2, age)

	age = 0b11001111    // assign as a binary number
	num4 = 0xffaabb1122 // hexa number

	println(age, num4)
	fmt.Printf("%b %b\n", age, num4)
	fmt.Printf("0x%x 0x%x\n", age, num4)
	fmt.Printf("%d %d\n", age, num4)

	// Group of variables
	var (
		a, b      = 10, 20
		f1, f2    = 1231.12312, 5454.454
		ok3, str3 = true, "Hello WOrld"
	)

	fmt.Println(a, b, ok3, str3)
	fmt.Printf("type of a: %T\n", a)
	fmt.Println("Type of f1", reflect.TypeOf(f1))
	fmt.Println("Type of f2", reflect.TypeOf(f2))

	// shorthand notation

	c, d := 10, 20
	// var c int = 10
	// var d int = 20
	str5, ok5 := "Hello Go Minds", true
	f3 := 12321.232

	c = 23213123123
	// c = "Hell World"

	fmt.Println(c, d, str5, ok5, f3)
}

// func Greet() {
// 	println(str1)
// }

// Zero value --> when the variable is declared but not assigned a value to it, it automatically assigned a default value

/*
Numbers:
int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64
uintptr
rune,byte (not a separate types but alias of int32 and uint8 respectively)


Bool:
bool --> true/false

String:
string --> string is collection of bytes

Any/Empty Interface:
any/interface{}

Complex:
	complex64,complex128
*/

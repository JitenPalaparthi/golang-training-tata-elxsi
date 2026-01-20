package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var Greet string = "Hello Tata Minds!"

func main() {

	println(Greet)
	greet := "Hello Tata Minds!"

	// Is string mutable?
	greet = "Hello Tata Folks!" // Mutate the string

	// こんにちは世界
	greet = "Hello ちは世界 💕" // utf-8

	greet = "Hello World! How are you doing! Hope you are travelling further!"

	fmt.Println(greet)
	fmt.Println("Type of:", reflect.TypeOf(greet), "size of:", unsafe.Sizeof(greet), "len of:", len(greet))

	//s1 := "Hello World" // Some mermory in DS
	//s2 := "Hello World" // Same ptr is used even for s2

	//var s3, s4, s5 string // ""

	// s1 := "Hello"
	// s2 := "World"
	// s3 := s1 + " " + s2 // CPU does some work

	var any1 any // emtpy interface

	// Is it type safe ? Yes
	// what is the zero value --> nil
	// what default type  --> nil

	any1 = true // What is the data and what type here

	any1 = greet

	any1 = 12312312.12312

	any1 = 12331231

	any1 = "Hello Tata!"

	println("size of any1:", unsafe.Sizeof(any1))
	fmt.Println(&greet)

	any1 = true
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	var any2 any
	fmt.Println("Value of any2:", any2, "Type of any2:", reflect.TypeOf(any2))

	if any2 == nil {
		println("any2 is nil, no data and no type")
	}

	c1 := 123.2 + 9.0i // r+i

	c2 := complex(123.23, 23)

	var r, i float32 = 123.32, 3.2

	c3 := complex(r, i) // complex64

	var c4 complex64

	c4 = complex(r, i)

	fmt.Println(c1, c2, c3, c4)

}

// any , interface{}
// string

// string --> String Header
// ------------------------
// DataPtr --> 8 bytes nil or some ptr
// Len --> 8 bytes

// any --> any structure
// DataPtr --> Pointer to the data
// TypePtr --> Pointer to the type

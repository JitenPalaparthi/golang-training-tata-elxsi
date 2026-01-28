package main

import "fmt"

func main() {

	// alias not new types
	var i integer = 100
	println(i)

	var ch char = 'A'
	println(string(ch))
	// using a new type

	var (
		mi1 myint1 = 12312

		mi2 myint2 = 23

		mi3 myint3 = 45

		i1 int = 10
	)

	str1 := mi1.ToString()
	println(str1)

	sq1 := myint2(mi1).Sq()

	println(sq1)

	cb1 := myint3(mi1).Cube()

	println(cb1)

	str2 := myint1(mi2).ToString()
	println(str2)

	sq2 := mi2.Sq()

	println(sq2)

	cb2 := myint3(mi2).Cube()

	println(cb2)

	str3 := myint1(mi3).ToString()
	println(str3)

	sq3 := myint2(mi3).Sq()

	println(sq3)

	cb3 := mi3.Cube()

	println(cb3)

	str4 := myint1(i1).ToString()
	println(str4)

	sq4 := myint2(i1).Sq()

	println(sq4)

	cb4 := myint3(i1).Cube()

	println(cb4)

}

type integer = int // alias of int , it is not a new type
type char = rune   // char = rune = int32

// type rune=int32
// type byte =uint8
// type any =interface{}

type myint1 int // myint is a new type

func (mi myint1) ToString() string {
	return fmt.Sprint(mi)
}

type myint2 int // myint is a new type

func (mi myint2) Sq() int {
	return int(mi * mi)
}

type myint3 int // myint is a new type

func (mi myint3) Cube() int {
	return int(mi * mi * mi)
}

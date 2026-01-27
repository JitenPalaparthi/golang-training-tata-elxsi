package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var p1 Person

	p1.Id = 123
	p1.Name = "Jiten"
	p1.Email = "JitenP@outlook.Com"
	p1.Mobile = "9618558500"

	fmt.Println(p1)

	p2 := Person{Id: 103, Name: "Jiten", Email: "JItenP@Outlook.Com", Mobile: "9618558500"}

	fmt.Println(p2)

	p3 := new(Person) // It allocates memory and creates a struct based on the size of all the variables

	(*p3).Id = 123 // No need to give * operator even it is a pointer
	(*p3).Name = "Jiten"
	p3.Email = "JitenP@outlook.Com"
	p3.Mobile = "9618558500"

	fmt.Println(p3, "Size of p3:", unsafe.Sizeof(*p3))

	var t1 T1

	fmt.Println(t1, unsafe.Sizeof(t1)) //11 bytes

	var t2 T2

	fmt.Println(t2, unsafe.Sizeof(t2))

	var t3 T3

	fmt.Println(t3, unsafe.Sizeof(t3))

}

// User defined types

type Person struct {
	Id     int    // 8
	Name   string // 16 // string struct {ptr,len}
	Email  string // 16 // string struct {ptr,len}
	Mobile string // 16 // string struct {ptr,len}
}

// What is the size of the structure

type T1 struct { /// There are zero values for the struct
	ok1 bool  // 1 --> 8 , padding of 7 bytes
	id  int   // 8 --> 8
	age uint8 // 1 --> 8// there is an alias for this which is byte
	ok2 bool  // 1 --> 0
}

type T2 struct { /// There are zero values for the struct
	id  int   // 8 --> 8
	ok1 bool  // 1 --> 8
	age uint8 // 1 --> 8// there is an alias for this which is byte
	ok2 bool  // 1 --> 0
	// ok3 bool
	// ok4 bool
	// ok5 bool
	// ok6 bool
	// ok7 bool
	// ok8 bool
}

// 1------- 8------ 87------

type T3 struct { /// There are zero values for the struct
	ok1 bool  // 1 --> 4 --> padding 3 bytes
	id  int32 // 4 --> 4 --> paddings 0 bytes
	age uint8 // 1 --> 4 --> padding 3 bytes
	ok2 bool  // 1 --> 0 --> adjusted in the above allocation , padding 2 bytes
	ok3 bool  // 1 --> 0 --> adjusted in the above allocation , padding 1 bytes
	ok4 bool  // 1 --> 0 --> adjusted in the above allocation , no padding
}

// 4--- 4--- 4321

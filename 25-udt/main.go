package main

import "fmt"

func main() {

	var c1 ColourCode

	c1.int = 123
	c1.string = "red"

	fmt.Println(c1)

	c2 := ColourCode{int: 423, string: "blue"}

	fmt.Println(c2)

	p1 := Person{Id: 9878, Name: "Jiten", Address: &Address{Line1: "Kesavadasapuram", City: "Trivandrum", Pincode: "695011"}}

	a1 := new(Address)
	a1.City = "Trivandrum"
	a1.Pincode = "695011"
	a1.Line1 = "Kesavadasapuram"

	p2 := Person{Id: 12321, Name: "Jiten", Address: a1}

	a2 := &Address{Line1: "Kesavadasapuram", City: "Trivandrum", Pincode: "695011"}

	p3 := &Person{Id: 12321, Name: "Jiten", Address: a2}

	a3 := Address{Line1: "Kesavadasapuram", City: "Trivandrum", Pincode: "695011"}

	p4 := &Person{Id: 12321, Name: "Jiten", Address: &a3}

	p5 := Person{}

	p5.Id = 34234
	p5.Name = "Jiten"

	// p5.Address.City = "Trivandrum"

	// p5.Address.Line1 = "Ullor"
	// p5.Address.Pincode = "695011"

	p5.Address = &Address{Line1: "Kesavadasapuram", City: "Trivandrum", Pincode: "695011"}

	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(p3)
	// fmt.Println(p4)
	// fmt.Println(p5)
	PrintPerson(p1)

	p1.PrintPerson()
	PrintPerson(p2)
	p2.PrintPerson()

	PrintPerson(*p3)
	(*p3).PrintPerson() // no need to do any deref even p3 is a pointer, go can understand

	PrintPerson(*p4)
	p4.PrintPerson()
	PrintPerson(p5)
	p5.PrintPerson()
}

// struct with anonymous fields
type ColourCode struct { // anonymous fields of the struct
	string
	int
}

type Address struct {
	Line1   string
	City    string
	Pincode string
}

type Person struct {
	Id      int
	Name    string
	Address *Address // can give both field name as the type, no issues in Golangv, composition
}

func PrintPerson(p Person) { // It is a function
	println()
	println("Calling a function as Person variable as an argument")
	println("Id:", p.Id)
	println("Name:", p.Name)
	println("Address:")
	if p.Address != nil {
		println("Line1:", p.Address.Line1)
		println("City:", p.Address.City)
		println("Pincode:", p.Address.Pincode)
	} else {
		println("No address")
	}
	println("-----------\n")
}

func (p Person) PrintPerson() { // It is a method, attached to Person type
	println()
	println("Calling a method on Person type")
	println("Id:", p.Id)
	println("Name:", p.Name)
	println("Address:")
	if p.Address != nil {
		println("Line1:", p.Address.Line1)
		println("City:", p.Address.City)
		println("Pincode:", p.Address.Pincode)
	} else {
		println("No address")
	}
	println("-----------\n")
}

package main

func main() {

	//a, b, c := 10, 20, 0

	a, b := 10, 20
	// t := b
	// b = a
	// a = t
	a, b = b, a // easy swap

	//a1, b1, c1 := 10, 20, 30

	//a1, b1, c1 = b1, c1, a1

	c := (100) + (a+b)*10 + a/b + (a*b)/2 + 100 + (20 + b) + (a * 10) + 200 // is it an atomic operation?

	println(c)

	d := c > 10 // compare , d is bool

	e := d && false // logical operation
	println(e)

}

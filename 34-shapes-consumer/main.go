package main

import (
	"fmt"

	"github.com/JitenPalaparthi/tata-elxsi-shapes/shapes/rect"
	"github.com/JitenPalaparthi/tata-elxsi-shapes/shapes/square"
)

func main() {

	r1 := rect.New(10.4, 11.34)

	a1, p1 := r1.Area(), r1.Perimeter()

	fmt.Printf("Area of Rect r1:%.2f\nPerimeter of Rect r1:%.2f\n", a1, p1)

	s1 := square.New(23.45)

	println()
	a2, p2 := s1.Area(), s1.Perimeter()
	fmt.Printf("Area of Square s1:%.2f\nPerimeter of Square s1:%.2f\n", a2, p2)

	//rand.IntN(100)

	rect.Greet()

}

// run the command to get package from internet .. which is a third party package
// go get github.com/JitenPalaparthi/tata-elxsi-shapes

// to vendor --> it creates a vendor directory and pulls  all the source code of pacakges from internet

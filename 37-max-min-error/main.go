package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	println(max(10, 20))
	fmt.Println(max(12.2, 12.1)) // takes both int and float bcz 1.18 onwards go supports generics

	var slice []int // slice is nil

	for i := 0; i < 10; i++ {
		slice = append(slice, rand.IntN(100)) // using append . when slice is nil ,and using append would initilize the slice and then appends it
	}

	fmt.Println(slice)
	mx, mi, err := GetMax(slice)

	fmt.Println(mx, mi, err)

	if mx, mi, err = GetMax([]int{}); err != nil {
		println(err.Error())
	} else {
		fmt.Println("max:", mx, "min:", mi)
	}

	if mx, mi, err = GetMax(nil); err != nil {
		println(err.Error())
	} else {
		fmt.Println("max:", mx, "min:", mi)
	}

	var slice2 []int
	if mx, mi, err = GetMax(slice2); err != nil {
		println(err.Error())
	} else {
		fmt.Println("max:", mx, "min:", mi)
	}

}

// What should be an error ?
// Generally errors.New , fmt.Errorf()
// can creat any time that implements the interface called error is a valid type to be used as error

func GetMax(slice []int) (int, int, error) {
	if slice == nil {
		return 0, 0, errors.New("nil slice")
	}
	if len(slice) == 0 {
		return 0, 0, fmt.Errorf("slice with no length")
	}
	mx := slice[0]
	mi := slice[0]

	for _, v := range slice {
		mx = max(v, mx)
		mi = min(v, mi)
	}

	return mx, mi, nil
}

// error , io.writer, io.reader

// Make sure to return max and ensure that a and b are numbers and a and b are of same type
func Max(a, b any) (any, error) {

	return nil, nil
}

// Make sure to return Min and ensure that a and b are numbers and a and b are of same type

func Min(a, b any) (any, error) {

	return nil, nil
}

package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {

	//var num int
	var arr1 [5]int
	var arr2 [2]int

	// What is the zero value
	// [0 0 0 0 0]

	// var arr3 [3]int
	// arr1 = ([5]int)arr3 // cannot type cast

	fmt.Println("arr1 type:", reflect.TypeOf(arr1))
	fmt.Println("arr2 type:", reflect.TypeOf(arr2))
	fmt.Println(arr1, arr2)

	arr1[0], arr1[1], arr1[2] = 10, 11, 12
	arr1[3] = 13
	arr1[4] = 14
	//arr1[5] = 100 // invalid argument: index 5 out of bounds [0:5] up to 5 but not 5

	// sum := 0

	// for _, v := range arr1 {
	// 	sum += v
	// }
	// println(sum)

	sum := SumOfArr(arr1)
	println(sum)

	//str := "hello World"

	// ptr: &123123
	// len:11 --> [0:11]
	sum = 0
	for i := 0; i < len(arr1); i++ { // length is taken from where? [5]int
		sum += arr1[i]
	}
	println(sum)

	//sum = SumOfArr2(arr2)

	// for i := 0; i < len(str); i++ { // // length is taken from where? Len:11 from string header
	// 	print(string(str[i]))
	// }

	arr3 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} // directly assign values
	arr4 := [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 123, 234, 234, 234, 23, 345, 65, 675, 3, 45}
	fmt.Println(arr3, arr4)

	for i := range arr3 {
		arr3[i] = rand.IntN(999)
	}

	fmt.Println(&arr3[0]) // 0x140000b8050

	//var arr5 [999999999999]int
	//	fmt.Println(&arr5[0])

}

// Array is a collection of elements of the same type
// Array is has a fixd length -> The length is known to the compiler, array size cannot be grown or shrunk
// What is the type of an array ? [5]int The type of the array includes its length [5]int
// cannot type cast between different types of arrays
// Is array mutable?

func SumOfArr(arr [5]int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOfArr2(arr [2]int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}
	return sum
}

func MaxOfArray(arr [5]int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func SunOf(any1 any) any {

	var sum int

	switch ar := any1.(type) {

	case [2]int:
		for _, v := range ar {
			sum += v
		}
	case [3]int:
		for _, v := range ar {
			sum += v
		}

	case [4]int:
		for _, v := range ar {
			sum += v
		}

	case [5]int:
		for _, v := range ar {
			sum += v
		}
	default:
		return nil
	}

	// for _, v := range any1 {

	// }

	return sum
}

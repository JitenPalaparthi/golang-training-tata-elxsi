package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15}

	slice2 := slice1 /// header copy

	println(slice2)

	slice3 := make([]int, 10)

	copy(slice3, slice1) // dest to src
	/// it is a deep copy means both slice3 and slice1 are different
	// copy copies elemtns based on the dest size ..
	// slice3 10,11,12,13,14,15,0,0,0,0

	fmt.Println(slice3)

	// clear

	clear(slice3) // keep the len but makes all values as zero values

	fmt.Println(slice3)

	map1 := make(map[string]any)
	map1["name"] = "Jiten"
	clear(map1) // make the map as no keys and values.. but map is not nil
	fmt.Println(map1)
}

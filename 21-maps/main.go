package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 100)

	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}

	fmt.Println(slice1)

	map1 := make(map[int]int, 100)

	for _, v := range slice1 {

		val, ok := map1[v]
		if ok {
			map1[v] = val + 1
		} else {
			map1[v] = 1
		}

		// val := map1[v]
		// map1[v] = val + 1
	}

	for k, v := range map1 {
		println("Key:", k, "Value:", v)
	}

}

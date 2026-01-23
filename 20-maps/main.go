package main

import (
	"errors"
	"fmt"
)

func main() {

	var map1 map[string]any
	// what datatype can be key, any data type on which we can apply == operator can be a key
	// this is only declaration of map but not instantiation

	if map1 == nil {
		fmt.Println("nil map")
		map1 = make(map[string]any, 20) // no size if given
		// map1 creates two buckets to begin with (it depends on the version and the internal algarthim)
		// each bucket comntains three parallel arrays with 8 elements
		// why 8 elements ? 8 is not too big neither too small, it was thought that it is a balanced length
	}

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560042"] = "Bangalore-4"
	map1["560093"] = 560093

	// the map, picks the key , example 560086
	// it creates hash 648a6a 6ffff -> modulus operation get a number --> bucket number <n umber of buckets
	// top hash 648
	// key 560086
	// "Bangalore-1"

	for key, val := range map1 {
		fmt.Println("Key:", key, "Value:", val)
	}

	val := map1["560093"]
	fmt.Println(val)

	val, ok := map1["31243432"] // same pattern like we have done with any, ok is bool , if true yes there is the key else key does not exist
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("key does not exist")
	}

	//delete(map1, "560093") // delete deletes if key exists, if does not exist , delete does noting

	if err := Delete(map1, "560093"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	}
	if err := Delete(map1, "5600wqeqw93"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	} // delete deletes if key exists, if does not exist , delete does noting
	//delete(map1, "560") // delete deletes if key exists, if does not exist , delete does noting

	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	map2 := map[string]string{"Hello": "Hi", "name": "Jiten"} // creating a map with direct values
	fmt.Println(map2)
	var map3 map[string]any

	err := Delete(map3, "Hello")
	if err != nil {
		println(err.Error())
	}

	//delete(map3, "hello")
}

func Delete(m map[string]any, key string) error {
	if m == nil {
		return fmt.Errorf("map is nil")
	}
	if _, ok := m[key]; !ok {
		return errors.New("key does not exist")
	}

	delete(m, key)
	return nil
}

// why not this way --> ?
func Delete1(m map[string]any, key string) string {
	if m == nil {
		return "map is nil"
	}
	if _, ok := m[key]; !ok {
		return "key does not exist"
	}

	delete(m, key)
	return ""
}

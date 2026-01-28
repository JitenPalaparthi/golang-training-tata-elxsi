package main

import "fmt"

func main() {

	mymap1 := make(mymap, 5)
	mymap1["name"] = "Jiten"
	mymap1["age"] = 41
	mymap1["address"] = struct {
		Line1   string
		PinCode string
		City    string
	}{Line1: "PRRA45", PinCode: "690511", City: "Trivandrum"}

	keys, values := mymap1.GetKeysNVals()
	fmt.Println(keys)
	fmt.Println(values)

}

type mymap map[string]any

func (mm mymap) GetKeys() []string {
	keys := make([]string, len(mm))
	i := 0
	for k := range mm {
		keys[i] = k
		i++
	}
	return keys
}

func (mm mymap) GetValues() []any {
	values := make([]any, len(mm))
	i := 0
	for _, v := range mm {
		values[i] = v
		i++
	}
	return values
}

func (mm mymap) GetKeysNVals() ([]string, []any) {
	keys := make([]string, len(mm))
	values := make([]any, len(mm))
	i := 0
	for k, v := range mm {
		keys[i] = k
		values[i] = v
		i++
	}
	return keys, values

}

// Write delete function

// create a delte method

// create a normal map and type caste to mymap and access those methods

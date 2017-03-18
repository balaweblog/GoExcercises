package main

import (
	"fmt"
)

func main() {

	//map
	// panic: assignment to entry in nil map
	var collections map[int]string
	collections = make(map[int]string)
	collections[0] = "first"
	collections[1] = "second"
	fmt.Println(collections)

	coll := map[int]string{
		1: "team",
		2: "test",
	}
	coll[3] = "seems"
	fmt.Println(coll)
	fmt.Println(len(coll))

	//update maps
	coll[1] = "onion"
	fmt.Println(coll)
	fmt.Println(len(coll))

	//delete maps
	delete(coll, 2)
	fmt.Println(coll)
	fmt.Println(len(coll))

	//check if the values exists
	// if the key exists return the value
	if val, exists := coll[3]; exists {
		fmt.Println(val)
		fmt.Println("value exists")
	} else {
		fmt.Println("not exists")
	}

	//search and delete
	if val, exists := coll[3]; exists {
		fmt.Println(exists, val)
		delete(coll, 3)
	} else {
		fmt.Println("not")
	}

	//for loop map
	for key, val := range coll {
		fmt.Println(key, val)
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")
	//User defined types
	// Go allows us the ability to declare our own types.
	//This is bad practice
	//conversion:int(myAge) and converting type foo to type int
	var age Foo
	age = 1
	fmt.Println(age)
	var calc int
	calc = int(age)
	fmt.Println(calc)
}

type Foo int

package main

import (
	"fmt"
)

func main() {

	var a int = 20
	var b int = 11

	fmt.Println(a * b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c float32 = 20.2
	var d float32 = 11.2

	//float you will get the output in float variables
	fmt.Println(c * d)
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c / d)
	// modulo wont work in floatfmt.Println(c % d)

	// this wont work since it is mismatched types int and float64
	// you need to convert appropriate
	/*
	    fmt.Println(a / c)
		fmt.Println(a * c)
		fmt.Println(a % c)

		fmt.Println(d / b)
		fmt.Println(d * b)
		fmt.Println(d % b)
	*/

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")

	//basic if else
	if true {
		fmt.Println("i am true")
	} else if false {
		fmt.Println("am in false")
	} else {
		fmt.Println("i am in betweens")
	}

	//not exclation
	a := 10
	if !(a > 20) {
		fmt.Println("<10")
	}

	//beauty of go initializatin
	//if we are going to use the field only in the if field then you can declare it.
	if ab := 10; true {
		fmt.Println(ab)
	}
}

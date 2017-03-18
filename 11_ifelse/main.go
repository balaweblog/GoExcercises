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
	//do some calculation if result<10 do something else do somethong
	if ab := 10; true {
		fmt.Println(ab)
	}

	//if else or and and
	//ab:=20
	cc := 22
	if ab := 20; ab == 20 && cc == 22 || false {
		fmt.Println("okd..")
	}
}

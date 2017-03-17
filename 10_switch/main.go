package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")
	i := 10
	// basic switch
	// no break needed
	//default is optional
	switch i {
	case 1:
		fmt.Println("i am in 1")
	case 2:
		fmt.Println("i am in 2")
	case 3:
		fmt.Println("i am in 3")
	default:
		fmt.Println("defaulter")
	}

	//fallthrough
	//fallthroufh cant be there for last statement
	//if any statement satisfy along with fallthrough then all the  cases after that will also execute
	//ie falling through from here one
	//if nothing statisfy default will go
	//Note: fallthrough will reach default as well irrespective of cases or default based on fallthrough
	j := 20
	switch j {
	case 10:
		fmt.Println("i am in 1")
		fallthrough
	case 20:
		fmt.Println("i am in 2")
		fallthrough
	case 30:
		fmt.Println("i am in 3")
		fallthrough
	case 40:
		fmt.Println("i am here in 4")
	default:
		fmt.Println("default")
	}

	//multiple cases values
	ab := 11
	switch ab {
	case 11, 12:
		fmt.Println("11 or 12")
	case 13, 14:
		fmt.Println("13 or 14")
	case 15, 16:
		fmt.Println("15 or 16")
	}

	//cases with expressions

	r := "test"
	m := 10
	switch {
	case r == "test" && m < 2:
		fmt.Println("am in new place")
	case r == "test" && m < 12:
		fmt.Println("am in second place")
	}

	// switch on types
	//interface can acccepte any values casting boxing/unboxing will work here
	switchontypes("test")
	switchontypes(1)
    var see float32
	switchontypes(see)

}
func switchontypes(types interface{}) {
	switch types.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float32:
		fmt.Println("float")
	}
}

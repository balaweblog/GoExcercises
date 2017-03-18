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
    
    //user defined types extension
    ppl := People{"balamurugan", "alagumalai", "dhiva"}
	fmt.Println(ppl.Length())
    ppl.OrderAsc()
    fmt.Println(ppl)
    
    
}

type Foo int


type People []string

func (p People) Length() int {
	return len(p)
}
func (p People) OrderAsc() {
	var c People
	for _, i := range p {
		for _, j := range p {
			if (i != j) && i > j {
				c = append(c, i)
			}
		}
	}
	p = c
}

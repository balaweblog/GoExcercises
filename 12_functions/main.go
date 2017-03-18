package main

import (
	"fmt"
)

func main() {
	// main is the entry point to your program
	// main is the first function.
	fmt.Println("hi")

	// pass in parameter function
	saysomething("hello")
	saysomething("how")

	// pass two parameters
	greetmessage("Welcome", "balamurugan")

	//return values
	fmt.Println(add(1, 3))

	//named returns avoid using it
	fmt.Println(sub(1, 3))

	//multiple return values
	fmt.Println(arith(1, 3))

	//variadic params
	addlist(1, 3, 3, 3, 33, 3)
	addlist(1, 1)

	//variadic variable and arguments
	data := []int{1, 3, 3, 3, 3, 2, 3, 3, 33}
	addlist(data...)

	//func expreessions as function
	team := func() {
		fmt.Println("hi")
	}

	team()

	//function expressions as types
	//if you see this the output will be a memory address in "org" ie you can defainte function as variable
	//variable func expression
	org := func() int {
		return 10
	}
	fmt.Println(org)   //wont work
	fmt.Println(org()) // will work

	//function can return as func()
	//remind how are you are calling. since it returns functions you need to call as function
	// or else it will return only the address
	value := values()
	fmt.Println(value())

	/*
	   {}
	   closure helps us limit the scope of variables used by multiple functions
	   without closure, for two or more funcs to have access to the same variable,
	   that variable would need to be package scope
	*/

	//call back function
	// callback: passing a func as an argument
	//while returning call this function
	//go to inner function and come back here and print the function
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]

	//recursive
	//function can call itself
	fmt.Println(factorial(4))

	//defer
	//defer statement defers the execution of current statement untill the surrounding function returns
	//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
	//the output here will be world, hello,hi. will wait for main function returns then executes from returns
	printworld()
	defer sayhi()
	defer printhello()

	//function pass by values
	sei := 10
	changeme(sei)
	fmt.Println(sei)

	//functions pass by reference ie pointer values
	sed := 10
	changemes(&sed)
	fmt.Println(sed)

	//passing reference types in functions
	//serves a purpose different from new(T). It creates arrays, slices, maps, and channels only,
	home := Home{}
	home.Name = "bala"
	fmt.Println(home)
	printhome(home)
	fmt.Println(home)

	m := make(map[string]int)
	m["test"] = 11
	fmt.Println(m)
	dothis(m)
	fmt.Println(m)

	res := []string{"test"}
	fmt.Println(res)
	changeres(res)
	fmt.Println(res)

	home1 := Home{}
	home1.Name = "first "
	fmt.Println(home1.Name)
	printhome1(&home1)
	fmt.Println(home1.Name)

	//self executing functions
	func() {
		fmt.Println("am in self executing")
	}() //() is key for self executing function

}
func printhome1(home *Home) {
	home.Name = "net"
	fmt.Println(*home)

}
func changeres(res []string) {
	fmt.Println(res)
	res[0] = "new"
	fmt.Println(res)
}
func dothis(m map[string]int) {
	m["test"] = 1
}

func printhome(home Home) {
	fmt.Println(home)
	home.Name = "murugan"
	fmt.Println(home)
}

type Home struct {
	Name string
}

func changemes(i *int) {
	fmt.Println(*i)
	*i = 23
	fmt.Println(*i)
}
func changeme(i int) {
	fmt.Println(i)
	i = 2
	fmt.Println(i)
}
func sayhi() {
	fmt.Println("hi")
}
func printhello() {
	fmt.Println("hello")
}
func printworld() {
	fmt.Println("world")
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func values() func() string {
	return func() string {
		return "test"
	}
}
func saysomething(value string) {
	fmt.Println(value)
}

func greetmessage(in string, out string) {
	fmt.Println(in, out)
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) (c int) {
	c = a - b
	return c
}

func arith(a, b int) (int, int, int, int) {
	return a + b, a - b, a / b, a % b
}
func addlist(af ...int) {
	for _, i := range af {
		fmt.Println(i)
	}
}

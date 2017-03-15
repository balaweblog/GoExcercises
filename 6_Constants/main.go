package main

import (
	"fmt"
)

const i int = 10

func main() {
	//simple unchanging value
	const str string = "test"
	fmt.Println(i)
	fmt.Println(str)

	//multiple intialization
	const (
		pie = 10
		sie = 2
		che = "test"
	)
	fmt.Println(pie)
	fmt.Println(sie)
	fmt.Println(che)

	//iota values
	const (
		a = iota
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//iota multiplication
	const (
		aa = iota
		ab = iota * 10
		ac = iota * 20
		_  = iota * 30
		ae = iota * 40
	)
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ae)

	//iota expressions
	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
		GB = 1 << (iota * 10) // 1 << (3 * 10)
		TB = 1 << (iota * 10) // 1 << (4 * 10)
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}

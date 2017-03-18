package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int to other types
	var i int
	i = 32
	fmt.Println(i)
	fmt.Println(float32(i))
	fmt.Println(strconv.Itoa(i))

	//float32 to other types
	var j float64
	j = 22.2
	fmt.Println(j)
	fmt.Println(int(j))
	fmt.Println(strconv.FormatFloat(j, 'E', -1, 64))

	//rune to string
	var k rune
	k = 'a'
	fmt.Println(k)
	fmt.Println(string(k))

	//string to rune
	var l string
	l = "test"
	fmt.Println(l)
	fmt.Println([]byte(l))

	//string to int
	fmt.Println(strconv.Atoi("2"))

	//byte array to string
	fmt.Println(string([]byte{'a', 'b', 'c'}))

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	e, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, e, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)

}

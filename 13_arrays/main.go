package main

import (
	"fmt"
)

func main() {
	//arrays
	var arr [10]int
	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	//loop in arrays
	var arr1 [4]float32
	arr1[0] = 2.2
	arr1[1] = 2.43
	arr1[2] = 4.4
	for _, i := range arr1 {
		fmt.Println(i)
	}

	//string arrays
	var arr2 [20]string
	arr2[2] = "test"
	fmt.Println(arr2)

	// rune arrays will store in as character equivalent number
	var arr3 [12]rune
	arr3[2] = 'A'
	fmt.Println(arr3)

	var arr4 [20]byte
    arr4[2] = byte(22)
	fmt.Println(arr4)
    
    //length auto
    arr5:=make([]int,22)
    arr5[2] =2
    fmt.Println(arr5)
}

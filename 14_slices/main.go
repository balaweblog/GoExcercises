package main

import (
	"fmt"
)

func main() {

	//slices like arrays
	slice := make([]string, 10)
	slice[1] = "test"
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	//slices with append
	// 10 is length - number of elements referred to by the slice
	// 10 is capacity - number of elements in the underlying array
	myslice := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		myslice = append(myslice, i)
	}
	fmt.Println(myslice)
	fmt.Println(len(myslice), cap(myslice))

	//slicing a slice
	halfslice := myslice[2:20]
	fmt.Println(halfslice)
	fmt.Println(myslice[2])

	test := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(test[:2])
	fmt.Println(test[2:])
	fmt.Println(test[:])
	fmt.Println("ssss"[1], "ssss"[2])

	// length is autom incremented and append will add to it irrespecitve of length initially.
	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	fmt.Println(len(greeting), cap(greeting))
	greeting = append(greeting, "Suprabadham")
	fmt.Println(greeting[3])
	fmt.Println(len(greeting), cap(greeting))

	//append beyond capacity the capacity will get doubled than you intalized now it will be 5 to 10
	greeting = append(greeting, "news")
	greeting = append(greeting, "india")
	fmt.Println(greeting)
	fmt.Println(len(greeting), cap(greeting))

	//delete slice
	//slicing further is delete
	greeting = greeting[0:2]
	fmt.Println(greeting)

	//multidimension array
	student := []string{}
	studens := [][]int{}
	fmt.Println(student)
	fmt.Println(studens)

}

package main

import (
	"fmt"
)

func main() {
	//rune
	letter := rune("A"[0])
	fmt.Println(letter)

	ls := rune("test"[2])
	fmt.Println(ls)

	//string index access
	word := "Hello"
	letter1 := rune(word[0])
	fmt.Println(letter1)

	//hash function
	n := hashBucket("Go", 12)
	fmt.Println(n)
}
func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

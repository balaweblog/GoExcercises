package main

import (
	"fmt"
)

func main() {
	// setup pipeline.
	//getting multiple input
	//do calculation and expert

	//get the input in stream for 5 values
	ch := getinput(2, 3, 4, 4, 4, 2, 1, 1, 1)
	out := square(ch)
	for i := range out {
		fmt.Println(i)
	}
	//do the calculation

}
func getinput(values ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range values {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(ch chan int) chan int {

	out := make(chan int)
	go func() {
		for i := range ch {
			val := i * i
			out <- val
		}
		close(out)
	}()
	return out
}

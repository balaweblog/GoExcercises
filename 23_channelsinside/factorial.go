package main

import (
	"fmt"
)

func main() {
	//write a factorial probelem
	fmt.Println(factorial(40))

	//just convert the problem in channels
	output := factorialchannel(40)
	for i := range output {
		fmt.Println(i)
	}

	//just factoiral for number of values in channel pipeline
	input := getinput(10, 20, 30, 40)
	outputs := facpipeline(input)
	for i := range outputs {
		fmt.Println(i)
	}
}
func factorial(i int) int {
	var total int = 1
	for count := i; count > 0; count-- {
		total *= count
	}

	return total
}

func factorialchannel(i int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for count := i; count > 0; count-- {
			total *= count
		}
		out <- total
		close(out)
	}()

	return out
}

func getinput(values ...int) chan int {
	done := make(chan int)
	go func() {
		for _, i := range values {
			done <- i
		}
		close(done)
	}()
	return done
}

func facpipeline(input chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range input {
			total := 1
			for count := i; count > 0; count-- {
				total *= count
			}
			out <- total

		}
		close(out)
	}()
	return out
}

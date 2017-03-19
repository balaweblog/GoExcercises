package main

import (
	"fmt"
)

func main() {
	/*
							//dead lock situation
						    //deadlock is a state in which each member of a group of actions, is waiting for some other member to release a lock.
						    //why?
							c := make(chan int)
							c <- 1
							fmt.Println(<-c)


					//solution for deadlock
					c := make(chan int)
					go func() {
						c <- 1
					}()
					fmt.Println(<-c)


			    //another dead lock siutation
			    //the output will be 0
			    // ideally the channel has all but printed only 0
			    //why?
				c := make(chan int)

				go func() {
					for i := 0; i < 10; i++ {
						c <- i
					}
				}()

				fmt.Println(<-c)


			//the solution for the above problem is
		    // this will print 0 to 9 and will throw all go routines are sleep deadlock
		    //why
			c := make(chan int)

			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
			}()

			for {
				fmt.Println(<-c)
			}
	*/

	//always remmeber to close the channel once you put your messages
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}

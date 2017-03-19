package main

import (
	"fmt"
	//"sync"
	//"time"
)

func main() {
	//unbuffered channels
	/*
							c := make(chan int)

							go func() {
								for i := 0; i < 10; i++ {
									c <- i
								}
							}()

							go func() {
								for {
									fmt.Println(<-c)
								}
							}()
							time.Sleep(time.Second)


					    // go range over in channels
						d := make(chan int)
						go func() {
							for i := 0; i < 10; i++ {
								d <- i
							}
					        close(d) //close the channel
						}()

						go func() {
							for i := range d {
								fmt.Println(i)
							}
						}()
					    time.Sleep(time.Second)

			    //having two channels
			    //semapshore - sharing resources
				ch := make(chan int)
				done := make(chan bool)

				go func() {
					for i := 0; i < 10; i++ {
						ch <- i
					}
					done <- true
				}()
				go func() {
					for i := 10; i < 20; i++ {
						ch <- i
					}
					done <- true
				}()

			    // we block here until done <- true rather it should call in seperate thread. the func will wait till both complete.
				//<-done
				//<-done
				//close(c)
				go func() {
					<-done
					<-done
					close(ch)
				}()

				for c := range ch {
					fmt.Println(c)
				}

		//pass channels in parameters and return values
		ch := addition()
		done := sum(ch)

		for c := range done {
			fmt.Println(c)
		}
	*/
    //simple channels
	do := make(chan bool)

	go func() {
		for i := 0; i < 20000; i++ {
			if i > 2000 {
				do <- true
			}
		}
		close(do)
	}()
	<-do
	fmt.Println("i")
}
func sum(ch chan int) chan int {

	done := make(chan int)
	var sum int
	go func() {
		for i := range ch {
			sum = sum + i
		}
		done <- sum
		close(done)
	}()

	return done
}
func addition() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

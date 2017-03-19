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

						// channel direction for send and receive
						//If no direction is given, the channel is bidirectional.
						channel := make(chan int)
						go func() {
							for i := 0; i < 10; i++ {
								channel <- i
							}
					        close(channel)
						}()

						final := cleaner(channel)
						for i := range final {
							fmt.Println(i)
						}

				//closure bindings has all meening
			    //one might think the answer is a,b,c but instead the answer is c,c,c
			    // this is because each iteration of loop uses the same instances of the variable v. ie closure shares the same variable v.
			    //these things can be caught by go vet.
				done := make(chan bool)

				values := []string{"a", "b", "c"}
				for _, v := range values {
					go func() {
						fmt.Println(v)
						done <- true
					}()
				}

				// wait for all goroutines to complete before exiting
				for _ = range values {
					<-done
				}

		//closure for inner loop. ie create variable for each iteration. one way is to pass the value as argument to the closure.
		done := make(chan bool)

		values := []string{"a", "b", "c"}
		for _, v := range values {
			go func(u string) {
				fmt.Println(u)
				done <- true
			}(v)
		}

		// wait for all goroutines to complete before exiting
		for _ = range values {
			<-done
		}
	*/

	//closure another binding
	//create a new variable and assign to it.
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}
func cleaner(ch <-chan int) <-chan int {
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

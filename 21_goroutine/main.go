package main

import (
	"fmt"
	//"runtime"
	"sync/atomic"

	"sync"
	"time"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup
var counter int64 = 0

func main() {
	/*
			//no concurrency
			printfirst()
			printsecond()

			//go concurrency or go routine
			//nothing will happen none of the output will come
			go printfirst()
			go printsecond()

			//using waitgroup
			//the first one will print and second one will print but end at once since it concurrency
			wg.Add(2) //two methods
			go printfirst1()
			go printsecond1()
			wg.Wait()

			//go max procs parallel thread
			runtime.GOMAXPROCS(runtime.NumCPU())

			//assume during wait there is time sleep for few minutes
			//the output will be mixed since there is wait and thread will move to the next one.
			wg1.Add(2)
			go printfifth()
			go printsixth()
			wg1.Wait()

		//race condition
		//A race condition is when two or more routines have access to the same resource, such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines. This type of code can create the craziest and most random bugs you have ever seen
		//The program spawns two routines that each increment the global Counter variable twice. When both routines are done running, the program displays the value of the global Counter variable.
		//go build -race
	    //without timer the answer will be 6 with timer the answer in counter will be 2
	    //without timer - race condition
	    //with timer - mutex ie lock

	    //mutext condition
	    //The mutex class is a synchronization primitive that can be used to protect shared data from being simultaneously accessed by multiple threads. mutex offers exclusive, non-recursive ownership semantics: A calling thread owns a mutex from the time that it successfully calls either lock or try_lock until it calls unlock
		for i := 0; i <= 2; i++ {
			wg.Add(1)
			go find(i)
		}
		wg.Wait()
		fmt.Println(counter)
	*/

	//atomicity
    //output be 0. counter will reset in that particular routine and wont return back anything
	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go finder(i)
	}
	wg.Wait()
	fmt.Println(counter)

}
func finder(id int) {
	for i := 0; i < 2; i++ {
		value := counter
		time.Sleep(1 * time.Nanosecond)
		atomic.AddInt64(&counter, 1)
		//value++
		counter = value
	}
	wg.Done()
}
func find(id int) {
	for i := 0; i < 2; i++ {
		value := counter
		time.Sleep(1 * time.Nanosecond)

		value++
		counter = value
	}
	wg.Done()
}

func printfifth() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Millisecond)
	}
	wg1.Done()
}
func printsixth() {
	for i := 51; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Millisecond)
	}
	wg1.Done()
}

func printfirst() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}
}
func printsecond() {
	for i := 51; i < 100; i++ {
		fmt.Println(i)
	}
}
func printfirst1() {
	for i := 0; i < 500; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func printsecond1() {
	for i := 500; i < 1000; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

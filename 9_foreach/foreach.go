package main

import (
	"fmt"
)

func main() {

	/*
	   //basic for
	   for i:=0;i<10;i++{
	       fmt.Println(i)
	   }

	   j:=0
	   //while like for loop
	   for j<10{
	       fmt.Println(j)
	       j++
	   }

	   //nested for
	   for i:=100;i<200;i++{
	       for j:=0;j<10;j++{
	           fmt.Println(i,j)
	       }
	   }

	   //empty for loop
	   k:=0
	   for{
	       fmt.Println("test")
	       fmt.Println(k)
	       k++
	   }
	*/

    //breaks here the for loop and come out
	l := 10
	for {
		if l > 20 {
			break
		}
		fmt.Println(l)
		l++
	}

    //continue stops the loop thread further and go to next set of values
	m := 40
	for {
		m++
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
		if m >= 50 {
			break
		}
	}
}

package main

import (
	"fmt"
)

func main() {
    /*
	var i string = "test"
	//value
	fmt.Println(i)
	//address
	fmt.Println(&i)
    
    //scans and get the input and store in this address
    var values float32
    fmt.Print("enter your Salary")
    fmt.Scan(&values)
    fmt.Println(values*12)

    
    ///will thse address changes anywhere lets see
    // a pointer variable var b *int can hold address of other variable
    //if it is pointer varibale *b
    var a int  =10
    var b *int
    fmt.Println(a,&a)
    b =&a
    fmt.Println(b,&b)
    a =3
    fmt.Println(a,&a)
    fmt.Println(b,&b,*b)
    
    //we can pass memory address instead of a bunch of values 
    // pass by reference.
    // pass only the address ie when we pass the memory address we are passing a value.
    *b = 100
    fmt.Println(b,*b)
    */
    
    //no pointer
    // pass by values no change in values and look at their address address differnet
    x:=100
    change(x)
    fmt.Println(x,&x)
    
    //pass by reference
    //same address and values changed
    y:=200
    dontchange(&y)
    fmt.Println(y,&y)
    
    

}
func change(x int){
    x=200
    fmt.Println(x,&x)
}

func dontchange(x *int){
    *x =100
    fmt.Println(x,*x)
}
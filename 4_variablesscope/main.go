package main

import(
"fmt")
 var z =0
var m =120
func main(){
    //variable scope
    x:=100
    fmt.Println(x)
     foo()
    
    //y scope is inside println cant access outside 
    t := 42
	fmt.Println(t)
	{
		fmt.Println(t)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
    
    //for two or more funcs to have access to the same variable that variable would need to be package scope
    fmt.Println(increment())
	fmt.Println(increment())
    
    //anonymous function or assigning a func to the variable.
    p:=0
    decrement :=func() int{
        p--
        return p
    }
    fmt.Println(decrement())
    fmt.Println(decrement())
    
    //package scope
    fmt.Println(m)
    deep()
    
    //order matters
    uv:= 100
    fmt.Println(uv)
    //fmt.Println(pv)
    //var pv =102
    
    //variable shadowing bad coding practice
    max:=max(7)
    fmt.Println(max)
    
    //same package but in different files
    // you cant use go run main.go since file is in different one 
    //use go build && 4_variablesscope.exe
    fmt.Println(ten)
    
    
    
}
func max(x int)int {
    return x*x
}
func deep(){
    fmt.Println(m)
}
func foo(){
    // cant access x inside another function. since x is accessible only by the particular function.
    //fmt.Println("am inside",x)
}

func increment() int {
	z++
	return z
}
package main

import "fmt"

func main() {

    /* shorthand variables*/
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
    
    /* var zero valued */
    var h int
	var i string
	var j float64
	var k bool

	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Println()
    
    
    /* variable declaration */
    var message string
	message = "Hello World."
	fmt.Println(message)
    
    /* declare many at once */
    var message0 string
	var l, m, n int
	l = 1

	message0 = "Hello World!"

	fmt.Println(message0, l,m,n)
    
    
    /* init many at once */
    var message1 = "Hello World!"
	var o,p,q int = 1, 2, 3

	fmt.Println(message1,o,p,q)
    
    /* type inference */
    var message2 = "Hello World!"
	var r,s,t = 1, 2, 3

	fmt.Println(message2, r,s,t)
    
    /* infer mixed type */
    
    var message3 = "Hello World!"
	var u,v,w= 1, false, 3

	fmt.Println(message3, u,v,w)
    
    
    /* init shorthand */
    // you can only do this inside a func
	x,y,z := 1, false, 3

	fmt.Println(x,y,z)
    
    
    
    
}
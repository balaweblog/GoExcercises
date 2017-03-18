package main

import (
	"fmt"
)

func main() {

	// polymorphism
	//interface
	circle := Circle{1}
	c := area(circle)
	fmt.Println(c)

	rect := Rectangle{20, 20}
	r := area(rect)
	fmt.Println(r)

	//empty interfaces
	shapes := []Shapes{circle, rect}
	fmt.Println(shapes)

	//parameters accepting interface
	calculator(10, 2000)

	//interface array ie slices of any array
	arr := []interface{}{shapes}
	fmt.Println(arr)

}
func calculator(a interface{}, b interface{}) {
	fmt.Println(a, b)
}

//empty interface
type Shapes interface{}

type Circle struct {
	radius int
}
type Rectangle struct {
	length  int
	breadth int
}
type Square struct {
	height int
}

func (c Circle) area() int {
	return c.radius * c.radius
}
func (r Rectangle) area() int {
	return 2 * r.length * r.breadth
}
func (s Square) area() int {
	return 4 * s.height
}

type Shape interface {
	area() int
}

func area(s Shape) int {
	return s.area()
}

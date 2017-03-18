package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//go object oriented programming
	//1. encapsulation
	//state(fields),behaviour(method),exported or non exported

	//can be used in same package though Id or id but not in other package
	person := Person{1, "bala", 1}
	//can be exported or non exported
	person.id = 2 //go fields ie is state of obeject
	fmt.Println(person)
	//go methods - behaviour of object
	fmt.Println(person.getName())

	//2. embedded types - reusablity - inheritance
	school := School{1, Person{2, "murugan", 33}}
	fmt.Println(school)

	//3. promotion - overriding
	// same field name in one class and another class.
	// fields and methods of the inner-type are promoted to the outer-type
	employee := Employee{"balamurugan", Person{11, "balamurugan", 31}}
	fmt.Println(employee)

	// struct pointer
	news := &Employee{"bala", Person{1, "te", 2}}
	fmt.Println(news)

	//marshalling and unmarshalling
	val, _ := json.Marshal(news)
	fmt.Println(val) //- marsallled byte array
	fmt.Println(string(val))

	team := []byte(`{1,'test',1}`)
	json.Unmarshal(team, person)
	fmt.Println(person)

	//encoding/decoding
	json.NewEncoder(os.Stdout).Encode(person)

	var p1 Person
	per := strings.NewReader(`{"id":2,"Name":"sees","Age":33}`)
	json.NewDecoder(per).Decode(&p1)
	fmt.Println(p1.Name)

}

type Person struct {
	id   int
	Name string
	Age  int
}

func (p Person) getName() string {
	return p.Name
}

type School struct {
	schoolid int
	person   Person
}

type Employee struct {
	Name   string
	person Person
}

package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	// var fred person
	bob := person{}
	julia := person{
		"julia",
		40,
		"cat",
	}
	beth := person{
		age:  30,
		name: "Beth",
	}
	bob.name = "Bob"
	fmt.Println(bob.name)
	fmt.Println(julia)
	fmt.Println(beth)
}

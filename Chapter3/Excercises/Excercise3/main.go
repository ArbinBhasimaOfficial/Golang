package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"Arbin", "Bhasima", 101}
	emp2 := Employee{
		firstName: "Zion",
		lastName:  "Sarashiki",
		id:        102,
	}

	var emp3 Employee
	emp3.firstName = "Zizi"
	emp3.lastName = "Manon"
	emp3.id = 103

	fmt.Println("Employee one: ", emp1)
	fmt.Println("Employee two: ", emp2)
	fmt.Println("Employee three: ", emp3)
}

package main

import "fmt"

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "zion"

	pet := struct {
		name string
		kind string
	}{
		name: "zion",
		kind: "dog",
	}

	fmt.Println(person)
	fmt.Println(pet)
}

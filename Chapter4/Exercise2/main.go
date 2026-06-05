package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNums := make([]int, 100)
	for i := range randNums {
		randNums[i] = rand.Intn(101)
	}
	fmt.Println(randNums)

	for _, value := range randNums {
		switch {
		case value%6 == 0:
			fmt.Println("Six!")
		case value%2 == 0:
			fmt.Println("Two!")
		case value%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

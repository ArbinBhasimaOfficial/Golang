package main

import (
	"fmt"
	"math/rand"
)

// func main() {
// 	var randNums []int
// 	for i := 0; i < 100; i++ {
// 		randNums = append(randNums, rand.Intn(101))
// 	}
// 	fmt.Println(randNums)
// }

// using for range
func main() {
	randNums := make([]int, 100)
	for i := range randNums {
		randNums[i] = rand.Intn(101)
	}
	fmt.Println(randNums)
}

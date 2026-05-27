package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)
	fmt.Println(xArrayPointer)
}

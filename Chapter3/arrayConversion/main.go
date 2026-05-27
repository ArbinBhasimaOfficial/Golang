package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
	panicArray := [5]int(xSlice)
	fmt.Println(panicArray)
}

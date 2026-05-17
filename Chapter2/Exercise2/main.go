package main

import "fmt"

func main() {
	const value = 300
	var i int
	var f float64
	i = int(value)
	f = float64(value)
	fmt.Println(i, f)
}

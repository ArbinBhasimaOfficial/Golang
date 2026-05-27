package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s, b, s2, s3, s4)
	var x string = "hello 😊"
	var x2 string = x[4:7]
	var x3 string = x[:5]
	var x4 string = x[6:]
	fmt.Println(x, x2, x3, x4)
	fmt.Println(len(x))
}

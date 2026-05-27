package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}
	subSlices1 := greetings[0:2]
	subSlices2 := greetings[1:4]
	subSlices3 := greetings[3:5]

	fmt.Println("Original Slice: ", greetings)
	fmt.Println("First Sub Slice Contains:", subSlices1)
	fmt.Println("Second Sub Slice Contains:", subSlices2)
	fmt.Println("Third Sub Slice Contains:", subSlices3)
}

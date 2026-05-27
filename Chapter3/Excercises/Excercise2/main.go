package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)

	fmt.Printf("The fourth rune is: %c\n", runes[3])
}

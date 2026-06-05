package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}
	for _, a := range words {
		switch wordLen := len(a); {
		case wordLen < 5:
			fmt.Println(a, "is short word!")
		case wordLen >= 10:
			fmt.Println(a, "is long word!")
		default:
			fmt.Println(a, "is exactly the right length")
		}

		switch {
		case len(a) == 2:
			fmt.Println("length is 2")
		case len(a) == 3:
			fmt.Println("length is 3")
		case len(a) == 4:
			fmt.Println("length is 4")
		default:
			fmt.Println("a is ", a)
		}

		switch len(a) {
		case 2:
			fmt.Println("length is 2")
		case 3:
			fmt.Println("length is 3")
		case 4:
			fmt.Println("length is 4")
		default:
			fmt.Println("a is ", a)
		}
	}
}

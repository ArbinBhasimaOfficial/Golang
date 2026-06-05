package main

import "fmt"

// func main() {
// 	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

// 	for _, word := range words {
// 		switch size := len(word); size {
// 		case 1, 2, 3, 4:
// 			fmt.Println(word, "is a short word")
// 		case 5:
// 			wordLen := len(word)
// 			fmt.Println(word, "is exactly the right lendgth:", wordLen)
// 		case 6, 7, 8, 9:
// 		default:
// 			fmt.Println(word, "is a long word")
// 		}
// 	}
// }

func main() {
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

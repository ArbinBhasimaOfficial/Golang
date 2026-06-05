// package main

// import "fmt"

// func main() {
// 	var total int

// 	for i := 0; i < 10; i++ {

// 		// total := total + i this to
// 		total = total + i // this
// 		fmt.Println(total)

// 	}
// 	fmt.Println(total)
// }

// in range variation
package main

import "fmt"

func main() {
	var total int

	// Create a slice with numbers 0 through 9
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Using for range instead of a standard counter loop
	for _, i := range numbers {
		// total := total + i // The same variable shadowing bug is here!
		total = total + i
		fmt.Println(total)
	}

	fmt.Println("Final total:", total)
}

// The bug in this code is Variable Shadowing caused
// by using the short declaration operator (:=) inside the loop body.

// When you wrote total := total + i,
// the := told Go to create a brand-new variable named total
// that only exists inside the scope of that specific loop iteration.

// Inside the loop: T
// his new, inner total shadows (hides)
// the original total declared at the top of main.
// It calculates 0 + i, prints it, and
// then is immediately destroyed when the loop iteration ends.

// Outside the loop:
// The original total variable declared at the top of main was
// never actually modified.
// That is why it prints Final total: 0 at the very end.

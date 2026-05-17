package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	t := []string{"A", "B", "C"}
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	// fmt.Println(slices.Equal(x, s)) // doesnot compile
	fmt.Println("EqualFunc(s, t):",
		slices.EqualFunc(s, t, func(a, b string) bool {
			return strings.EqualFold(a, b)
		}),
	)
	fmt.Println("Compare(x, y):", slices.Compare(x, y)) // 0
	fmt.Println("Compare(x, z):", slices.Compare(x, z)) // -1
	fmt.Println("Compare(z, x):", slices.Compare(z, x)) // 1
	fmt.Println("CompareFunc(s, t):",
		slices.CompareFunc(s, t, func(a, b string) int {
			return strings.Compare(
				strings.ToLower(a),
				strings.ToLower(b),
			)
		}),
	)
}

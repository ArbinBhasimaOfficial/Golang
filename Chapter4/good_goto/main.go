package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated studd no matter whe we left the loop")
	fmt.Println(a)
}

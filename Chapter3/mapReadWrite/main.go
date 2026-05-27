package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["kittens"])
	totalWins["kittens"]++
	fmt.Println(totalWins["kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
}

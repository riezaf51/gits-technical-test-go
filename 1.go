package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)
	if input <= 0 {
		fmt.Println("Invalid input")
		return
	}

	var result []int = lazyCatererSequence(input)
	for i := 0; i < input; i++ {
		fmt.Print(result[i])
		if i < input-1 {
			fmt.Print("-")
		}
	}
}

func lazyCatererFormula(n int) int {
	return n*(n+1)/2 + 1
}

func lazyCatererSequence(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, lazyCatererFormula(i))
	}
	return result
}

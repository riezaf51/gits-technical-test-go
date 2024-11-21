package main

import (
	"fmt"
)

func main() {
	var numOfScores, numOfGitsScores int
	var scores, gitsScores []int

	fmt.Scan(&numOfScores)
	if numOfScores <= 0 {
		fmt.Println("Invalid input")
		return
	}
	scores = make([]int, numOfScores)
	for i := 0; i < numOfScores; i++ {
		fmt.Scan(&scores[i])
		if scores[i] < 0 {
			fmt.Println("Invalid input")
			return
		}
	}

	fmt.Scan(&numOfGitsScores)
	if numOfGitsScores <= 0 {
		fmt.Println("Invalid input")
		return
	}
	gitsScores = make([]int, numOfGitsScores)
	for i := 0; i < numOfGitsScores; i++ {
		fmt.Scan(&gitsScores[i])
		if gitsScores[i] < 0 {
			fmt.Println("Invalid input")
			return
		}
	}

	var result []int = denseRanking(numOfScores, scores, numOfGitsScores, gitsScores)

	for i := 0; i < numOfGitsScores; i++ {
		fmt.Print(result[i], " ")
	}
}

func denseRanking(numOfScores int, scores []int, numOfGitsScores int, gitsScores []int) []int {
	var result []int
	for i := 0; i < numOfGitsScores; i++ {
		var rank int = 1
		for j := 0; j < numOfScores; j++ {
			if scores[j] <= gitsScores[i] {
				break
			}
			if j == 0 || scores[j] != scores[j-1] {
				rank++
			}
		}
		result = append(result, rank)
	}
	return result
}

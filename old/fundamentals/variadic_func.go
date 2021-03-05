package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinishes(13, 5, 12, 74, 24, 87)
	fmt.Println(bestFinish)
}

// Passing an unknown number of parameters to a function
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

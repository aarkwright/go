package main

import (
	"fmt"
)

func main() {
	// Variables to store the course rankins
	if firstRank, secondRank := 3922, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" +
			" than the second course.")
	} else if firstRank > secondRank {
		fmt.Println("\nSecond course is doing better" +
			" than the first one.")
	} else {
		fmt.Println("\nBoth courses are either the same" +
			" or something werid is going on.")
	}
}

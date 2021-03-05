package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{"Docker", "RHEL6", "RHEL7"}
	coursesCompleted := []string{"Docker", "RHEL6"}

	for _, i := range coursesInProg {
		//fmt.Println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\t> Found a clash:", i, "is in both lists!")
			}
		}
	}
}

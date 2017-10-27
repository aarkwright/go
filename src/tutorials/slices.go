package main

import (
	"fmt"
)

func main() {

	// Declares a slice called myCourses

	//myCourses := make([]string, 5, 10)
	myCourses := []string{"Docker", "RHEL 6", "RHEL 7"}

	fmt.Printf("Length is: %d\nCapacity is: %d",
		len(myCourses), cap(myCourses))
}

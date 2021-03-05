package main

import "fmt"

func main() {
	
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	// or get the pointer to one
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "Nigel",
		Level: "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
}
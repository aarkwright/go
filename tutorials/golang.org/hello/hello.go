package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Setup the logger
	log.SetPrefix("greetings: ") // simple prefix
	log.SetFlags(0)              // disable printing time, source file and line number

	// A slice of names
	names := []string{
		"Apollo",
		"Scytale",
		"Giedi",
	}

	// Request a greeting message.
	message, err := greetings.Hellos(names)

	// Handle error if exists
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	fmt.Println(message)
}

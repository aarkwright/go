package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	tvshow := "Poirot"

	fmt.Println("\nHi", name, ", you're currenty watching", tvshow)

	changeTVShow(&tvshow)
	fmt.Println("\nYou've changed the TV channel, and you're now watching", tvshow, ".")
}

func changeTVShow(_tvshow *string) string {
	*_tvshow = "Narcos"

	fmt.Println("Trying to switch TV channel to", *_tvshow)

	return *_tvshow
}
package main

import "fmt"

func main() {
	
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for k, v := range testMap {
		fmt.Printf("\n[+] Key:\t%v\n[+] Value:\t%v", k, v)
	}
}
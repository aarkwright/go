package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.println("Hello from", runtime.GOOS)
}
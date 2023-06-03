package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run helloworld.go <name>")
		return
	}

	name := os.Args[1]
	fmt.Println("Hello", name, "!")
}

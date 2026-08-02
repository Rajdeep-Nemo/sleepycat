package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	name, err := sleepycat.String(
		sleepycat.Prompt("Enter username: "),
	)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	fmt.Printf("Username: %s\n", name)
}

package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	temp, err := sleepycat.Float32(
		sleepycat.Prompt("Enter temperature: "),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read temperature: %v", err)
	}
	fmt.Printf("Temperature entered: %.2f\n", temp)
}

package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	// Example 1: Basic integer input
	age, err := sleepycat.Int(
		sleepycat.Prompt("Enter your age: "),
	)
	if err != nil {
		log.Fatalf("Failed to read age: %v", err)
	}
	fmt.Printf("Your age is %d.\n\n", age)

	// Example 2: Retrying up to 3 times on invalid input
	pin, err := sleepycat.Int(
		sleepycat.Prompt("Enter a number (max 3 attempts): "),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed after 3 attempts: %v", err)
	}
	fmt.Printf("Number accepted: %d\n\n", pin)

	// Example 3: Infinite retries until a valid integer is provided
	score, err := sleepycat.Int(
		sleepycat.Prompt("Enter game score (retries until valid): "),
		sleepycat.MaxAttempt(0),
	)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Final Score recorded: %d\n", score)
}

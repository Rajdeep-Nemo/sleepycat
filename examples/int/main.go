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
	fmt.Printf("Final Score recorded: %d\n\n", score)

	// Example 4: Integer input with a minimum length (e.g. a 4+ digit PIN)
	pinCode, err := sleepycat.Int(
		sleepycat.Prompt("Enter a PIN (min 4 digits): "),
		sleepycat.MinLength(4),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read PIN: %v", err)
	}
	fmt.Printf("PIN accepted: %d\n\n", pinCode)

	// Example 5: Integer input with min and max length (e.g. a 4-6 digit code)
	otp, err := sleepycat.Int(
		sleepycat.Prompt("Enter a 4-6 digit OTP: "),
		sleepycat.MinLength(4),
		sleepycat.MaxLength(6),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read OTP: %v", err)
	}
	fmt.Printf("OTP accepted: %d\n", otp)
}

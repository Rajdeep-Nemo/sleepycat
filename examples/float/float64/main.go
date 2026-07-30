package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	price, err := sleepycat.Float64(
		sleepycat.Prompt("Enter price: "),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read price: %v", err)
	}
	fmt.Printf("Price entered: %.2f\n", price)
}

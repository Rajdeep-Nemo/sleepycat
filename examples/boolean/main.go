package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	enabled, err := sleepycat.Bool(
		sleepycat.Prompt("Enable feature? "),
		sleepycat.MaxAttempt(3),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enabled:", enabled)
}

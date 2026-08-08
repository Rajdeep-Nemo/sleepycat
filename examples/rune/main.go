package main

import (
	"fmt"
	"os"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	val, err := sleepycat.Rune(
		sleepycat.Prompt("Enter a character: "),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(val)
}

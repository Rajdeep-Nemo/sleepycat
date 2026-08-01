# sleepycat

`sleepycat` is a Go package for reading user input from the terminal.

The package provides typed input functions and a consistent option-based API for prompts, validation, and retry behavior.

> **Status:** Work in progress. The API is not yet stable.

## Installation

```bash
go get github.com/Rajdeep-Nemo/sleepycat
```

## Example

```go
package main

import (
	"fmt"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	age, err := sleepycat.Int(
		sleepycat.Prompt("Age: "),
		sleepycat.MaxAttempts(3),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(age)
}
```

## Available

### Input

* `Int()`
* `Float32()`
* `Float64()`

### Options

* `Prompt()`
* `MaxAttempt()`

## Project Structure

```text
sleepycat/
├── .github/
│   └── workflows/
│       └── ci.yml              # Continuous integration
│
├── examples/                   # Example programs
│   ├── int/
│   │   └── main.go             
│   └── ...                     # More Examples
│
├── internal/
│   └── input.go                # Terminal input helpers
│
├── config.go                   # Config and functional options
├── parser.go                   # Shared parsing/retry logic
├── int.go                      # Integer input
├── float32.go                  # 32 bit float input
├── float64.go                  # 64 bit float input
│
├── helpers_test.go             # Shared helper functions
├── int_test.go
├── float32_test.go
├── float64_test.go
│
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── CHANGELOG.md
└── README.md
```

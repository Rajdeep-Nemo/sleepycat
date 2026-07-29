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

### Options

* `Prompt()`
* `MaxAttempt()`
* `MinLength()`
* `MaxLength()`

## Project Structure

```text
sleepycat/
├── .github/
│   └── workflows/
│       └── ci.yml              # Continuous integration
│
├── examples/
│   └── int/
│       └── main.go             # Example program
│
├── internal/
│   └── input.go                # Terminal input helpers
│
├── config.go                   # Config and functional options
├── parser.go                   # Shared parsing/retry logic
├── int.go                      # Integer input
│
├── int_test.go
│
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── CHANGELOG.md
└── README.md
```

### Responsibilities

| File                   | Responsibility                                                                            |
| ---------------------- | ----------------------------------------------------------------------------------------- |
| `config.go`            | Defines the internal configuration, functional options, and default values.               |
| `parser.go`            | Implements the shared input loop (prompt → read → parse → retry).                         |
| `int.go`               | Public `Int()` input function.                                                            |
| `internal/input.go`    | Reads raw input from an `io.Reader`. Used internally by the package.                      |
| `examples/`            | Runnable examples demonstrating the public API.                                           |
| `*_test.go`            | Unit tests for each public input type.                                                    |


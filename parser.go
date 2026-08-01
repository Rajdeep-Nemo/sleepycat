package sleepycat

import (
	"fmt"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

// Checks for valid option values and take a parsing method,
// uses that method to parse the input string following options
func input[T any](parse func(string) (T, error), opts ...option) (T, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	var zero T

	if cfg.maxAttempt < 0 {
		return zero, fmt.Errorf("sleepycat: MaxAttempt must be >= 0, got %d", cfg.maxAttempt)
	}

	attempts := 0
	for {
		text, err := internal.Read(cfg.prompt)
		if err != nil {
			return zero, err
		}
		value, err := parse(text)
		if err == nil {
			return value, nil
		}
		attempts += 1
		if cfg.maxAttempt != 0 && attempts >= cfg.maxAttempt {
			return zero, err
		}
	}
}

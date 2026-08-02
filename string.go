package sleepycat

import (
	"strings"
)

// String reads from the input source and returns a string literal
// MaxAttempt does not work as there is no invalid string
func String(opts ...option) (string, error) {
	return input(func(text string) (string, error) {
		return strings.TrimSpace(text), nil
	}, opts...)
}

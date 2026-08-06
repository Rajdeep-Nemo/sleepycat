package sleepycat

import (
	"fmt"
)

// Parses the string value
func parseByte(s string) (byte, error) {
	if len(s) != 1 {
		return 0, fmt.Errorf("sleepycat: parsing %q: expected a single ASCII character", s)
	}
	return s[0], nil
}

// Read from the input source and returns a `byte` value
func Byte(opts ...option) (byte, error) {
	return input(parseByte, opts...)
}

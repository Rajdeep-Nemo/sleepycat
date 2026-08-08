package sleepycat

import (
	"fmt"
)

// Parses the string value
func parseRune(s string) (rune, error) {
	r := []rune(s)
	if len(r) != 1 {
		return 0, fmt.Errorf("sleepycat: parsing %q: expected a single character", s)
	}
	return r[0], nil
}

// Read from the input source and returns a unicode code point
func Rune(opts ...option) (rune, error) {
	return input(parseRune, opts...)
}

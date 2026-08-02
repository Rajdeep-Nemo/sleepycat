package sleepycat

import (
	"strconv"
)

// Bool reads from the input source and returns a boolean value (true/false)
func Bool(opts ...option) (bool, error) {
	return input(strconv.ParseBool, opts...)
}

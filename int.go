package sleepycat

import (
	"strconv"
)

// Int reads an integer from the input source.
func Int(opts ...option) (int, error) {
	return input(strconv.Atoi, opts...)
}

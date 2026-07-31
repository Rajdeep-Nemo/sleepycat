package sleepycat

import (
	"strconv"
)

// Float32 reads an floating point number from the input source and returns as `float32`
func Float32(opts ...option) (float32, error) {
	return input(func(s string) (float32, error) {
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	}, opts...)
}

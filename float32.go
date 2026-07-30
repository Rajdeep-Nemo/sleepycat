package sleepycat

import (
	"strconv"
)

func Float32(opts ...option) (float32, error) {
	return input(func(s string) (float32, error) {
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	}, opts...)
}

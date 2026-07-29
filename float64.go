package sleepycat

import (
	"strconv"
)

func Float64(opts ...option) (float64, error) {
	return input(func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}, opts...)
}

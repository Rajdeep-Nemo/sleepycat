package sleepycat

import (
	"strconv"
)

func Int(opts ...option) (int, error) {
	return input(strconv.Atoi, opts...)
}

package sleepycat

import (
	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func input[T any](parse func(string) (T, error), opts ...option) (T, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	// Loops input prompt until attempts are left
	attempts := 0
	for {
		var zero T

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

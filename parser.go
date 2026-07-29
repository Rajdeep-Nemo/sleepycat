package sleepycat

import (
	"fmt"
	"strings"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func checkLength(text string, cfg config) error {
	length := len([]rune(strings.TrimSpace(text)))
	switch {
	case length < cfg.minLength:
		return fmt.Errorf("min %d chars", cfg.minLength)
	case cfg.maxLength != -1 && length > cfg.maxLength:
		return fmt.Errorf("max %d chars", cfg.maxLength)
	}
	return nil
}

func input[T any](parse func(string) (T, error), opts ...option) (T, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	attempts := 0
	for {
		var zero T
		text, err := internal.Read(cfg.prompt)
		if err != nil {
			return zero, err
		}

		if lenErr := checkLength(text, cfg); lenErr != nil {
			attempts += 1
			if cfg.maxAttempt != 0 && attempts >= cfg.maxAttempt {
				return zero, lenErr
			}
			continue
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

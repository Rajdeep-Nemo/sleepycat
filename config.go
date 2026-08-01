package sleepycat

// Holds option values
type config struct {
	prompt     string // Input prompt
	maxAttempt int    // Maximum attempt before proper input
}

// Option configures the behavior of an input function.
type option func(*config)

// Default configuration to be used if no option is provided
func defaultConfig() config {
	return config{
		prompt:     "",
		maxAttempt: 1,
	}
}

// Prompt sets the text printed before waiting for input.
func Prompt(inputPrompt string) option {
	return func(c *config) {
		c.prompt = inputPrompt
	}
}

// MaxAttempts limits the number of invalid input attempts.
// A value of 0 retries indefinitely.
func MaxAttempt(count int) option {
	return func(c *config) {
		c.maxAttempt = count
	}
}

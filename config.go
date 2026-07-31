package sleepycat

// Holds option values
type config struct {
	prompt     string // Input prompt
	maxAttempt int    // Maximum attempt before proper input
	mask       rune   // Mask the input with provided symbol
	minLength  int    // Minimum length
	maxLength  int    // Maximum length
}

// Option configures the behavior of an input function.
type option func(*config)

// Default configuration to be used if no option is provided
func defaultConfig() config {
	return config{
		prompt:     "",
		maxAttempt: 1,
		minLength:  1,
		maxLength:  -1,
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

// MinLength determines the minimum length required.
// Setting it to 0 means empty input is allowed
func MinLength(min int) option {
	return func(c *config) {
		c.minLength = min
	}
}

// MaxLength determines the maximum number allowed.
func MaxLength(max int) option {
	return func(c *config) {
		c.maxLength = max
	}
}

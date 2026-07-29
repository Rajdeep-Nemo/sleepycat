package sleepycat

type config struct {
	prompt     string // Input prompt
	maxAttempt int    // Maximum attempt before proper input
	mask       rune   // Mask the input with provided symbol
	minLimit   int    // Minimum length
	maxLimit   int    // Maximum length
}

type option func(*config)

func defaultConfig() config {
	return config{
		prompt:     "",
		maxAttempt: 1,
	}
}

func Prompt(inputPrompt string) option {
	return func(c *config) {
		c.prompt = inputPrompt
	}
}

func MaxAttempt(count int) option {
	return func(c *config) {
		c.maxAttempt = count
	}
}

func MinLimit(min int) option {
	return func(c *config) {
		c.minLimit = min
	}
}

func MaxLimit(max int) option {
	return func(c *config) {
		c.minLimit = max
	}
}

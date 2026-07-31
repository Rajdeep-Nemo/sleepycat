package sleepycat

type config struct {
	prompt     string // Input prompt
	maxAttempt int    // Maximum attempt before proper input
	mask       rune   // Mask the input with provided symbol
	minLength  int    // Minimum length
	maxLength  int    // Maximum length
}

type option func(*config)

func defaultConfig() config {
	return config{
		prompt:     "",
		maxAttempt: 1,
		minLength:  1,
		maxLength:  -1,
	}
}

func Prompt(inputPrompt string) option {
	return func(c *config) {
		c.prompt = inputPrompt
	}
}

func MaxAttempt(count int) option {
	return func(c *config) {
		if count < 0 {
			count = 1
		}
		c.maxAttempt = count
	}
}

func MinLength(min int) option {
	return func(c *config) {
		c.minLength = min
	}
}

func MaxLength(max int) option {
	return func(c *config) {
		c.maxLength = max
	}
}

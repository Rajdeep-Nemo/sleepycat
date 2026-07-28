package sleepycat

type config struct {
	prompt     string
	maxAttempt int
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

package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestInt_ValidInput(t *testing.T) {
	withInput(t, "42")

	got, err := Int()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestInt_NegativeNumber(t *testing.T) {
	withInput(t, "-7")

	got, err := Int()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -7 {
		t.Errorf("expected -7, got %d", got)
	}
}

func TestInt_TrimsWhitespace(t *testing.T) {
	withInput(t, "   15   ")

	got, err := Int()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 15 {
		t.Errorf("expected 15, got %d", got)
	}
}

func TestInt_InvalidInput_DefaultMaxAttemptOne(t *testing.T) {
	withInput(t, "not-a-number")

	got, err := Int()
	if err == nil {
		t.Fatalf("expected an error, got nil (value=%d)", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt_EmptyInput_FailsToParse(t *testing.T) {
	// Not using Required(), so empty input should simply fail parsing
	// (strconv.Atoi("") returns an error) rather than being treated specially.
	withInput(t, "")

	got, err := Int()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt_RetriesUntilValid(t *testing.T) {
	// First two lines are invalid, third is valid. MaxAttempt(3) should
	// allow exactly enough tries to reach the good value.
	withInput(t, "abc", "xyz", "9")

	got, err := Int(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 9 {
		t.Errorf("expected 9, got %d", got)
	}
}

func TestInt_ExhaustsMaxAttempts(t *testing.T) {
	// All three lines are invalid; MaxAttempt(3) should try exactly 3
	// times and then return the last error.
	withInput(t, "a", "b", "c")

	got, err := Int(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt_MaxAttemptCountIsExact(t *testing.T) {
	// Regression test for the earlier off-by-one bug: MaxAttempt(2) must
	// make exactly 2 attempts, not 1. We feed exactly 2 bad lines; if Int
	// only tries once, the second bad line is left unread but the test
	// still can't tell from the outside — so instead we feed a GOOD value
	// as the 2nd line and confirm it's actually reached.
	withInput(t, "bad", "77")

	got, err := Int(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 77 {
		t.Errorf("expected 77 from the 2nd attempt, got %d", got)
	}
}

func TestInt_MaxAttemptZeroMeansInfinite(t *testing.T) {
	// Feed several bad lines followed by a good one; MaxAttempt(0) (or
	// omitted, if default were 0) should keep retrying indefinitely
	// until it succeeds.
	withInput(t, "x", "y", "z", "w", "100")

	got, err := Int(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestInt_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	// internal.Read treats EOF-with-content as a successful read.
	// Simulate that directly by feeding a reader with no trailing newline.
	internal.SetInput(strings.NewReader("55"))
	t.Cleanup(internal.ResetInput)

	got, err := Int()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 55 {
		t.Errorf("expected 55, got %d", got)
	}
}

func TestInt_EOFWithNoContentReturnsError(t *testing.T) {
	// Genuinely empty reader: Read should surface io.EOF (or similar)
	// as an unrecoverable error, not loop forever.
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Int()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%d", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %d", got)
	}
}

func TestInt_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "21")

	got, err := Int(Prompt("Enter your age: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 21 {
		t.Errorf("expected 21, got %d", got)
	}
}

func TestInt_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "12.5") // valid float, invalid int

	_, err := Int()
	if err == nil {
		t.Fatal("expected an error for a float given to Int()")
	}
	if !strings.Contains(err.Error(), "12.5") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}

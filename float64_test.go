package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestFloat64_ValidInput(t *testing.T) {
	withInput(t, "3.14")

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 3.14 {
		t.Errorf("expected 3.14, got %v", got)
	}
}

func TestFloat64_NegativeNumber(t *testing.T) {
	withInput(t, "-7.5")

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != -7.5 {
		t.Errorf("expected -7.5, got %v", got)
	}
}

func TestFloat64_IntegerLikeInput(t *testing.T) {
	// A plain integer string like "42" is valid float64 input too.
	withInput(t, "42")

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 42.0 {
		t.Errorf("expected 42.0, got %v", got)
	}
}

func TestFloat64_TrimsWhitespace(t *testing.T) {
	withInput(t, "   1.5   ")

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 1.5 {
		t.Errorf("expected 1.5, got %v", got)
	}
}

func TestFloat64_ScientificNotation(t *testing.T) {
	// strconv.ParseFloat64 accepts scientific notation; confirm it passes through.
	withInput(t, "1.5e3")

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 1500 {
		t.Errorf("expected 1500, got %v", got)
	}
}

func TestFloat64_InvalidInput_DefaultMaxAttemptOne(t *testing.T) {
	withInput(t, "not-a-float")

	got, err := Float64()
	if err == nil {
		t.Fatalf("expected an error, got nil (value=%v)", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat64_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "xyz", "2.71")

	got, err := Float64(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 2.71 {
		t.Errorf("expected 2.71, got %v", got)
	}
}

func TestFloat64_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "a", "b", "c")

	got, err := Float64(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat64_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "7.7")

	got, err := Float64(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != 7.7 {
		t.Errorf("expected 7.7 from the 2nd attempt, got %v", got)
	}
}

func TestFloat64_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "y", "z", "w", "9.9")

	got, err := Float64(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != 9.9 {
		t.Errorf("expected 9.9, got %v", got)
	}
}

func TestFloat64_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("5.5"))
	t.Cleanup(internal.ResetInput)

	got, err := Float64()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 5.5 {
		t.Errorf("expected 5.5, got %v", got)
	}
}

func TestFloat64_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Float64()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat64_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "12.34")

	got, err := Float64(Prompt("Enter price: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 12.34 {
		t.Errorf("expected 12.34, got %v", got)
	}
}

func TestFloat64_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "12.3.4") // malformed float

	_, err := Float64()
	if err == nil {
		t.Fatal("expected an error for malformed float input")
	}
	if !strings.Contains(err.Error(), "12.3.4") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}

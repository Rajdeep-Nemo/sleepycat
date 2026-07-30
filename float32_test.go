package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestFloat32_ValidInput(t *testing.T) {
	withInput(t, "3.14")

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != float32(3.14) {
		t.Errorf("expected 3.14, got %v", got)
	}
}

func TestFloat32_NegativeNumber(t *testing.T) {
	withInput(t, "-7.5")

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != float32(-7.5) {
		t.Errorf("expected -7.5, got %v", got)
	}
}

func TestFloat32_IntegerLikeInput(t *testing.T) {
	withInput(t, "42")

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 42.0 {
		t.Errorf("expected 42.0, got %v", got)
	}
}

func TestFloat32_TrimsWhitespace(t *testing.T) {
	withInput(t, "   1.5   ")

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 1.5 {
		t.Errorf("expected 1.5, got %v", got)
	}
}

func TestFloat32_InvalidInput_DefaultMaxAttemptOne(t *testing.T) {
	withInput(t, "not-a-float")

	got, err := Float32()
	if err == nil {
		t.Fatalf("expected an error, got nil (value=%v)", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat32_RetriesUntilValid(t *testing.T) {
	withInput(t, "abc", "xyz", "2.71")

	got, err := Float32(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != float32(2.71) {
		t.Errorf("expected 2.71, got %v", got)
	}
}

func TestFloat32_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "a", "b", "c")

	got, err := Float32(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat32_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "7.7")

	got, err := Float32(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != float32(7.7) {
		t.Errorf("expected 7.7 from the 2nd attempt, got %v", got)
	}
}

func TestFloat32_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "y", "z", "w", "9.9")

	got, err := Float32(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != float32(9.9) {
		t.Errorf("expected 9.9, got %v", got)
	}
}

func TestFloat32_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("5.5"))
	t.Cleanup(internal.ResetInput)

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != 5.5 {
		t.Errorf("expected 5.5, got %v", got)
	}
}

func TestFloat32_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Float32()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestFloat32_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "12.3.4") // malformed float

	_, err := Float32()
	if err == nil {
		t.Fatal("expected an error for malformed float32 input")
	}
	if !strings.Contains(err.Error(), "12.3.4") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}

func TestFloat32_PrecisionWithinFloat32Range(t *testing.T) {
	// A value with more precision than float32 can hold should still
	// parse successfully, just rounded to float32 precision.
	withInput(t, "0.1")

	got, err := Float32()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	want := float32(0.1)
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestFloat32_MinLength_TooShortThenValid(t *testing.T) {
	withInput(t, "1", "1.25")

	got, err := Float32(MinLength(4), MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected no error after retry, got %v", err)
	}
	if got != float32(1.25) {
		t.Errorf("expected 1.25, got %v", got)
	}
}

func TestFloat32_MaxLength_TooLongThenValid(t *testing.T) {
	withInput(t, "3.14159", "3.1")

	got, err := Float32(MaxLength(4), MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected no error after retry, got %v", err)
	}
	if got != float32(3.1) {
		t.Errorf("expected 3.1, got %v", got)
	}
}

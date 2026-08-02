package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestBool_ValidInput_True(t *testing.T) {
	withInput(t, "true")

	got, err := Bool()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_ValidInput_False(t *testing.T) {
	withInput(t, "false")

	got, err := Bool()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != false {
		t.Errorf("expected false, got %v", got)
	}
}

func TestBool_AcceptsShortForms(t *testing.T) {
	// strconv.ParseBool accepts 1, t, T, TRUE, true, True and 0, f, F, FALSE, false, False.
	cases := []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"t", true},
		{"T", true},
		{"TRUE", true},
		{"True", true},
		{"0", false},
		{"f", false},
		{"F", false},
		{"FALSE", false},
		{"False", false},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			withInput(t, tc.input)

			got, err := Bool()
			if err != nil {
				t.Fatalf("expected no error for %q, got %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("expected %v for %q, got %v", tc.want, tc.input, got)
			}
		})
	}
}

func TestBool_TrimsWhitespace(t *testing.T) {
	withInput(t, "   true   ")

	got, err := Bool()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_InvalidInput_DefaultMaxAttemptOne(t *testing.T) {
	withInput(t, "not-a-bool")

	got, err := Bool()
	if err == nil {
		t.Fatalf("expected an error, got nil (value=%v)", got)
	}
	if got != false {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestBool_EmptyInput_FailsToParse(t *testing.T) {
	withInput(t, "")

	got, err := Bool()
	if err == nil {
		t.Fatalf("expected an error for empty input, got value=%v", got)
	}
	if got != false {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestBool_YesNoNotAccepted(t *testing.T) {
	// strconv.ParseBool does not accept "yes"/"no" — confirms Bool() is
	// strict to Go's bool grammar rather than a looser yes/no interpretation.
	withInput(t, "yes")

	got, err := Bool()
	if err == nil {
		t.Fatalf("expected an error for \"yes\", got value=%v", got)
	}
	if got != false {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestBool_RetriesUntilValid(t *testing.T) {
	withInput(t, "nah", "maybe", "true")

	got, err := Bool(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "a", "b", "c")

	got, err := Bool(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%v", got)
	}
	if got != false {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestBool_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "bad", "false")

	got, err := Bool(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected the 2nd attempt to succeed, got error: %v", err)
	}
	if got != false {
		t.Errorf("expected false from the 2nd attempt, got %v", got)
	}
}

func TestBool_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "x", "y", "z", "w", "true")

	got, err := Bool(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success with infinite retries, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("true"))
	t.Cleanup(internal.ResetInput)

	got, err := Bool()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Bool()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%v", got)
	}
	if got != false {
		t.Errorf("expected zero value on failure, got %v", got)
	}
}

func TestBool_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "true")

	got, err := Bool(Prompt("Continue? "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != true {
		t.Errorf("expected true, got %v", got)
	}
}

func TestBool_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "maybe")

	_, err := Bool()
	if err == nil {
		t.Fatal("expected an error for an invalid bool string")
	}
	if !strings.Contains(err.Error(), "maybe") {
		t.Errorf("expected underlying strconv error to mention the bad input, got: %v", err)
	}
}
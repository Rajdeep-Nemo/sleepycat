package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestString_ValidInput(t *testing.T) {
	withInput(t, "hello")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "hello" {
		t.Errorf("expected \"hello\", got %q", got)
	}
}

func TestString_TrimsLeadingTrailingWhitespace(t *testing.T) {
	withInput(t, "   hello world   ")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "hello world" {
		t.Errorf("expected \"hello world\", got %q", got)
	}
}

func TestString_PreservesInternalWhitespace(t *testing.T) {
	withInput(t, "  hello   world  ")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "hello   world" {
		t.Errorf("expected internal spacing preserved, got %q", got)
	}
}

func TestString_EmptyInput_NeverFailsToParse(t *testing.T) {
	// String's parse function always succeeds (no length constraints,
	// no conversion that can fail) — empty input is a valid string.
	withInput(t, "")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error for empty input, got %v", err)
	}
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestString_WhitespaceOnlyInput_TrimsToEmpty(t *testing.T) {
	withInput(t, "     ")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "" {
		t.Errorf("expected empty string after trimming whitespace-only input, got %q", got)
	}
}

func TestString_MaxAttemptOptionHasNoEffectSinceParseNeverFails(t *testing.T) {
	// Because String's parse function can't fail, MaxAttempt should never
	// actually trigger a retry — the first read always succeeds.
	withInput(t, "only one line needed")

	got, err := String(MaxAttempt(1))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "only one line needed" {
		t.Errorf("expected \"only one line needed\", got %q", got)
	}
}

func TestString_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("no newline"))
	t.Cleanup(internal.ResetInput)

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error on EOF-terminated input, got %v", err)
	}
	if got != "no newline" {
		t.Errorf("expected \"no newline\", got %q", got)
	}
}

func TestString_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := String()
	if err == nil {
		t.Fatalf("expected an error on empty reader, got value=%q", got)
	}
	if got != "" {
		t.Errorf("expected zero value on failure, got %q", got)
	}
}

func TestString_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "Alice")

	got, err := String(Prompt("Enter your name: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "Alice" {
		t.Errorf("expected \"Alice\", got %q", got)
	}
}

func TestString_ContainsSpecialCharacters(t *testing.T) {
	withInput(t, "hello@world.com #tag $100")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "hello@world.com #tag $100" {
		t.Errorf("expected special characters preserved, got %q", got)
	}
}

func TestString_UnicodeInput(t *testing.T) {
	withInput(t, "héllo wörld 你好")

	got, err := String()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "héllo wörld 你好" {
		t.Errorf("expected unicode preserved, got %q", got)
	}
}

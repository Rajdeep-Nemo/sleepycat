package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestRune_ValidInput(t *testing.T) {
	withInput(t, "A")

	got, err := Rune()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 'A' {
		t.Errorf("expected %q, got %q", 'A', got)
	}
}

func TestRune_UnicodeInput(t *testing.T) {
	withInput(t, "é")

	got, err := Rune()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 'é' {
		t.Errorf("expected %q, got %q", 'é', got)
	}
}

func TestRune_EmojiInput(t *testing.T) {
	withInput(t, "😊")

	got, err := Rune()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != '😊' {
		t.Errorf("expected %q, got %q", '😊', got)
	}
}

func TestRune_TrimsWhitespace(t *testing.T) {
	withInput(t, "   中   ")

	got, err := Rune()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != '中' {
		t.Errorf("expected %q, got %q", '中', got)
	}
}

func TestRune_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Rune()
	if err == nil {
		t.Fatalf("expected an error, got value=%q", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %q", got)
	}
}

func TestRune_MultipleCharacters(t *testing.T) {
	withInput(t, "AB")

	got, err := Rune()
	if err == nil {
		t.Fatalf("expected an error, got value=%q", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %q", got)
	}
}

func TestRune_MultipleUnicodeCharacters(t *testing.T) {
	withInput(t, "😊中")

	got, err := Rune()
	if err == nil {
		t.Fatalf("expected an error, got value=%q", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %q", got)
	}
}

func TestRune_RetriesUntilValid(t *testing.T) {
	withInput(t, "AB", "😊中", "Ω")

	got, err := Rune(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != 'Ω' {
		t.Errorf("expected %q, got %q", 'Ω', got)
	}
}

func TestRune_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "AB", "😊中", "")

	got, err := Rune(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%q", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %q", got)
	}
}

func TestRune_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "AB", "Ω")

	got, err := Rune(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected second attempt to succeed, got %v", err)
	}
	if got != 'Ω' {
		t.Errorf("expected %q, got %q", 'Ω', got)
	}
}

func TestRune_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "AB", "😊中", "", "Ж")

	got, err := Rune(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success, got %v", err)
	}
	if got != 'Ж' {
		t.Errorf("expected %q, got %q", 'Ж', got)
	}
}

func TestRune_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("Ω"))
	t.Cleanup(internal.ResetInput)

	got, err := Rune()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != 'Ω' {
		t.Errorf("expected %q, got %q", 'Ω', got)
	}
}

func TestRune_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Rune()
	if err == nil {
		t.Fatalf("expected an error, got value=%q", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %q", got)
	}
}

func TestRune_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "中")

	got, err := Rune(Prompt("Enter a character: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != '中' {
		t.Errorf("expected %q, got %q", '中', got)
	}
}

func TestRune_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "AB")

	_, err := Rune()
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "single character") {
		t.Errorf("unexpected error: %v", err)
	}
}

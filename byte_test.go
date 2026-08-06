package sleepycat

import (
	"strings"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

func TestByte_ValidInput(t *testing.T) {
	withInput(t, "A")

	got, err := Byte()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != byte('A') {
		t.Errorf("expected %q, got %q", 'A', got)
	}
}

func TestByte_DigitInput(t *testing.T) {
	withInput(t, "7")

	got, err := Byte()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != byte('7') {
		t.Errorf("expected %q, got %q", '7', got)
	}
}

func TestByte_TrimsWhitespace(t *testing.T) {
	withInput(t, "   A   ")

	got, err := Byte()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != byte('A') {
		t.Errorf("expected %q, got %q", 'A', got)
	}
}

func TestByte_EmptyInput(t *testing.T) {
	withInput(t, "")

	got, err := Byte()
	if err == nil {
		t.Fatalf("expected an error, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestByte_MultipleCharacters(t *testing.T) {
	withInput(t, "AB")

	got, err := Byte()
	if err == nil {
		t.Fatalf("expected an error, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestByte_UnicodeCharacter(t *testing.T) {
	withInput(t, "😊")

	got, err := Byte()
	if err == nil {
		t.Fatalf("expected an error, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestByte_RetriesUntilValid(t *testing.T) {
	withInput(t, "AB", "😊", "Z")

	got, err := Byte(MaxAttempt(3))
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if got != byte('Z') {
		t.Errorf("expected %q, got %q", 'Z', got)
	}
}

func TestByte_ExhaustsMaxAttempts(t *testing.T) {
	withInput(t, "AB", "😊", "")

	got, err := Byte(MaxAttempt(3))
	if err == nil {
		t.Fatalf("expected an error after exhausting attempts, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestByte_MaxAttemptCountIsExact(t *testing.T) {
	withInput(t, "AB", "X")

	got, err := Byte(MaxAttempt(2))
	if err != nil {
		t.Fatalf("expected second attempt to succeed, got %v", err)
	}
	if got != byte('X') {
		t.Errorf("expected %q, got %q", 'X', got)
	}
}

func TestByte_MaxAttemptZeroMeansInfinite(t *testing.T) {
	withInput(t, "AB", "😊", "", "K")

	got, err := Byte(MaxAttempt(0))
	if err != nil {
		t.Fatalf("expected eventual success, got %v", err)
	}
	if got != byte('K') {
		t.Errorf("expected %q, got %q", 'K', got)
	}
}

func TestByte_EOFWithoutTrailingNewlineStillParses(t *testing.T) {
	internal.SetInput(strings.NewReader("Q"))
	t.Cleanup(internal.ResetInput)

	got, err := Byte()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != byte('Q') {
		t.Errorf("expected %q, got %q", 'Q', got)
	}
}

func TestByte_EOFWithNoContentReturnsError(t *testing.T) {
	internal.SetInput(strings.NewReader(""))
	t.Cleanup(internal.ResetInput)

	got, err := Byte()
	if err == nil {
		t.Fatalf("expected an error, got value=%v", got)
	}
	if got != 0 {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestByte_PromptOptionDoesNotAffectParsing(t *testing.T) {
	withInput(t, "M")

	got, err := Byte(Prompt("Enter a character: "))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != byte('M') {
		t.Errorf("expected %q, got %q", 'M', got)
	}
}

func TestByte_ReturnsUnderlyingParseError(t *testing.T) {
	withInput(t, "AB")

	_, err := Byte()
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "single ASCII character") {
		t.Errorf("unexpected error: %v", err)
	}
}

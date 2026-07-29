package sleepycat

import (
	"io"
	"testing"

	"github.com/Rajdeep-Nemo/sleepycat/internal"
)

type lineReader struct {
	lines []string
	pos   int
}

func (l *lineReader) Read(p []byte) (int, error) {
	if l.pos >= len(l.lines) {
		return 0, io.EOF
	}
	line := l.lines[l.pos] + "\n"
	n := copy(p, line)
	if n < len(line) {
		l.lines[l.pos] = line[n:]
		return n, nil
	}
	l.pos++
	return n, nil
}

func withInput(t *testing.T, lines ...string) {
	t.Helper()
	internal.SetInput(&lineReader{lines: lines})
	t.Cleanup(internal.ResetInput)
}

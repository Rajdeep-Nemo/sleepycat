package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Store the bufio.Reader globally so buffer state is preserved across reads
var reader = bufio.NewReader(os.Stdin)

// ** SetInput and ResetInput functions are meant for test cases only **
// Changes the input source.
func SetInput(r io.Reader) {
	if r != nil {
		reader = bufio.NewReader(r)
	}
}

// Default input source.
func ResetInput() {
	reader = bufio.NewReader(os.Stdin)
}

// Prints prompt and reads the input
func Read(prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}

	text, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF || text == "" {
			return "", err
		}
	}

	return strings.TrimSpace(text), nil
}

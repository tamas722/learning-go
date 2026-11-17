package scanning

import (
	"bufio"
	"io"
	"unicode"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// counter reads a text and returns the counted values.
func counter(reader io.Reader) int {
	// INSERT YOUR CODE HERE
	bufReader := bufio.NewReader(reader)
	count := 0

	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			break
		}
		if unicode.IsLower(r) {
			count++
		}
	}

	return count
}

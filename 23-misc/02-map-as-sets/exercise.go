package search

import (
	"io"
	"regexp"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// search reads a text and a word and returns true if the word appears in the text and false if it does not.
func contain(reader io.Reader, word string) bool {
	// INSERT YOUR CODE HERE
	data, err := io.ReadAll(reader)
	if err != nil {
		return false
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	cleaned := re.ReplaceAllString(string(data), " ")

	counts := make(map[string]int)
	for _, w := range strings.Fields(cleaned) {
		counts[strings.ToLower(w)]++
	}

	target := strings.ToLower(strings.TrimSpace(word))
	return counts[target] > 1
}

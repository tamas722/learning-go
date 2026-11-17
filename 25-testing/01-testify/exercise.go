package testimony

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func StrNCpyINbounds(t *testing.T) {
	dest := "Hello"
	strncpy(&dest, "Go", 2)
	assert.Equal(t, "Go", dest)
}

func StrNCpyEmptyDestination(t *testing.T) {
	dest := ""
	strncpy(&dest, "World", 3)
	assert.Equal(t, "Wor", dest)
}

func StrNCpyEmptySource(t *testing.T) {
	dest := "Hi"
	strncpy(&dest, "", 1)
	assert.Equal(t, "", dest)
}

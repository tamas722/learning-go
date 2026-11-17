package closures

import "errors"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func proxy(limit int, fn func() int) func() (int, error) {
	// INSERT YOUR CODE HERE
	count := 0
	return func() (int, error) {
		if count < limit {
			count++
			return fn(), nil
		}
		return 0, errors.New("call limit exceeded")
	}
}

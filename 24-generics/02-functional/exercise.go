package functional

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func reduce[E, R any](slice []E, initial R, fn func(R, E) R) R {
	acc := initial
	for _, elem := range slice {
		acc = fn(acc, elem)
	}
	return acc
}

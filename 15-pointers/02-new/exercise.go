package pointernew

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func newValue() *float64 {
	// INSERT YOUR CODE HERE
	p := new(float64)
	*p = 12.05
	return p
}

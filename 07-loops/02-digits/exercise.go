package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func multiplyDigits(n int) int {
	if n == 0 {
		return 0
	}

	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}
	return product
}

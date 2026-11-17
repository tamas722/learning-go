package calculator

import (
	"math"
	"strconv"
)


// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE



func gmean(x, y float64) int {
	return int(math.Round(math.Sqrt(x * y)))
}

func gmeanString(x, y string) (int, error) {
	x1, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, err
	}
	y1, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, err
	}
	return gmean(x1, y1), nil
}
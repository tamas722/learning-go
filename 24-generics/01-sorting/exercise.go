package sorting

import "sort"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func sortSlice[T string | uint8](input []T) []T {
	result := make([]T, len(input))
	copy(result, input)

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

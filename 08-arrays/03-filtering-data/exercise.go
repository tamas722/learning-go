package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	// INSERT YOUR CODE HERE
	var result [10]string

	if len(keys) != len(indices) {
		return result
	}

	pos := 0
	for i, idx := range indices {
		if idx%2 != 0 {
			result[pos] = keys[i]
			pos++
			if pos == 10 {
				break
			}
		}
	}

	return result
}

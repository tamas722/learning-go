package sleepSort

import "time"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// reverseSleepSort returns the input uint-slice sorted in the reverse order.
func reverseSleepSort(input []uint) []uint {
	if len(input) == 0 {
		return []uint{}
	}

	ch := make(chan uint, len(input))

	for _, x := range input {
		go func(val uint) {
			time.Sleep(time.Duration(500-val*10) * time.Millisecond)
			ch <- val
		}(x)
	}

	result := make([]uint, 0, len(input))
	for i := 0; i < len(input); i++ {
		result = append(result, <-ch)
	}

	return result
}

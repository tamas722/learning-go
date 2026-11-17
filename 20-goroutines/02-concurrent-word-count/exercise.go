package wordcount

import (
	"strings"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func CountWords(lines []string) map[string]int {
	results := make(chan map[string]int)
	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			localCounts := make(map[string]int)
			for _, word := range strings.Fields(text) {
				localCounts[word]++
			}
			results <- localCounts
		}(line)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalCounts := make(map[string]int)
	for local := range results {
		for word, count := range local {
			totalCounts[word] += count
		}
	}

	return totalCounts
}

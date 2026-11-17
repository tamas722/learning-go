package channelmultiplexer

import (
	"context"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	out := make(chan any)
	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(c <-chan any) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-c:
					if !ok {
						return
					}
					select {
					case out <- val:
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

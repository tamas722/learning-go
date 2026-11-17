package channelbroadcaster

import "context"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {
	go func() {
		defer func() {
			for _, out := range outputs {
				if out != nil {
					close(out)
				}
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				for _, out := range outputs {
					if out == nil {
						continue
					}
					select {
					case out <- val:
					default:
						// Ha a kimeneti csatorna tele van, nem blokkolunk (5. követelmény)
					}
				}
			}
		}
	}()
}

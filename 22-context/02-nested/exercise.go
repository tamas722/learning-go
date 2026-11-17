package subtask

import (
	"context"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

type subTaskResult struct {
	res string
	err error
}

func StartTask(ctx context.Context) (result string, err error) {
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resCh := make(chan subTaskResult, 1)

	go func() {
		defer close(resCh)
		r, e := SubTask(subCtx)
		resCh <- subTaskResult{res: r, err: e}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case out := <-resCh:
		if out.err != nil {
			return "", out.err
		}
		return "Main task status: " + out.res, nil
	}
}

func SubTask(ctx context.Context) (result string, err error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return "Subtask completed successfully", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

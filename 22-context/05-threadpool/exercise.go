package threadpool

import (
	"context"
	"fmt"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

type Runnable interface {
	Run(context.Context) error
}

type ThreadPool interface {
	Run(Runnable)
	Close()
}

type threadPool struct {
	tasks     chan Runnable
	errCh     chan error
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	closeOnce sync.Once
}

func NewThreadPool(n int) (ThreadPool, chan error) {
	if n <= 0 {
		n = 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	tasks := make(chan Runnable, n*10)
	errCh := make(chan error, n*10)

	tp := &threadPool{
		tasks:  tasks,
		errCh:  errCh,
		ctx:    ctx,
		cancel: cancel,
	}

	for i := 0; i < n; i++ {
		tp.wg.Add(1)
		go tp.worker()
	}

	return tp, errCh
}

func (tp *threadPool) worker() {
	defer tp.wg.Done()

	for {
		select {
		case <-tp.ctx.Done():
			return
		case task, ok := <-tp.tasks:
			if !ok || task == nil {
				return
			}
			if err := task.Run(tp.ctx); err != nil {
				select {
				case tp.errCh <- err:
				default:
					fmt.Println("error channel full, dropping error:", err)
				}
			}
		}
	}
}

func (tp *threadPool) Run(task Runnable) {
	if task == nil {
		return
	}

	select {
	case <-tp.ctx.Done():
		return
	case tp.tasks <- task:
	}
}

func (tp *threadPool) Close() {
	tp.closeOnce.Do(func() {
		tp.cancel()
		tp.wg.Wait()
		close(tp.errCh)
	})
}

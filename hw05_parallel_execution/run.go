package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	var errCnt int32
	wg := sync.WaitGroup{}
	wg.Add(n)
	tasckCh := make(chan Task)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range tasckCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}
		tasckCh <- task
	}
	close(tasckCh)

	wg.Wait()
	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

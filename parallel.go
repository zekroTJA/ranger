package ranger

import (
	"context"
	"sync"
)

// Job defines a function that is taken in a list in Parallel which executes
// a task that takes some amount of time and returns a value T.
type Job[T any] func(ctx context.Context) T

// Parallel takes a list ob jobs and executes them in order using a pool of
// goroutines. The poolSize defines the amount of created goroutines. So this
// is the amount of tasks which are executed at the same time at maximum.
//
// The jobs result values are yielded for the loop.
func Parallel[T any](jobs []Job[T], poolSize int) func(yield func(T) bool) {
	if poolSize < 1 {
		poolSize = 1
	}

	return func(yield func(T) bool) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(poolSize)

		queue := make(chan Job[T], poolSize)
		defer close(queue)

		for range poolSize {
			go func() {
				defer wg.Done()

				for job := range queue {
					res := job(ctx)
					select {
					case <-ctx.Done():
						return
					default:
						if !yield(res) {
							cancel()
						}
					}
				}
			}()
		}

		for _, job := range jobs {
			select {
			case <-ctx.Done():
				return
			default:
				queue <- job
			}
		}

		wg.Wait()
	}
}

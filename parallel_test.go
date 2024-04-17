package ranger

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func ExampleParallel() {
	jobs := []Job[int]{
		func(_ context.Context) int {
			time.Sleep(200 * time.Millisecond)
			return 1
		},
		func(_ context.Context) int {
			time.Sleep(100 * time.Millisecond)
			return 2
		},
		func(_ context.Context) int {
			time.Sleep(200 * time.Millisecond)
			return 3
		},
		func(_ context.Context) int {
			time.Sleep(100 * time.Millisecond)
			return 4
		},
		func(_ context.Context) int {
			time.Sleep(200 * time.Millisecond)
			return 5
		},
	}

	for r := range Parallel(jobs, 2) {
		fmt.Println(r)
	}

	// Output:
	// 2
	// 1
	// 4
	// 3
	// 5
}

func TestParallel_break(t *testing.T) {
	jobs := []Job[int]{
		func(_ context.Context) int {
			time.Sleep(200 * time.Millisecond)
			println("foo", 1)
			return 1
		},
		func(_ context.Context) int {
			time.Sleep(100 * time.Millisecond)
			println("foo", 2)
			return 2
		},
		func(_ context.Context) int {
			time.Sleep(200 * time.Millisecond)
			println("foo", 3)
			return 3
		},
	}

	for v := range Parallel(jobs, 2) {
		println(v)
		break
	}
}

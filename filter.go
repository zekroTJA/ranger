package ranger

func Filter[T any](s []T, predicate func(i int, v T) bool) func(yield func(i int, v T) bool) {
	return func(yield func(i int, v T) bool) {
		for i, v := range s {
			if !predicate(i, v) {
				continue
			}

			if !yield(i, v) {
				break
			}
		}
	}
}

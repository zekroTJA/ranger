package ranger

import "fmt"

func ExampleFilter() {

	s := []int{1, 2, 3, 4, 5, 6}
	pret := func(i, v int) bool {
		return v%2 == 0
	}

	for i, v := range Filter(s, pret) {
		fmt.Printf("%d: %d\n", i, v)
	}

	// Output:
	// 1: 2
	// 3: 4
	// 5: 6
}

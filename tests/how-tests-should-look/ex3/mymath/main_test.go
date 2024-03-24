package mymath

import (
	"fmt"
	"testing"
)

// Examples for the CenteredAvg function
func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// Output:
	// 5
}

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	// composite literal table for test struct
	tests := []test{
		test{[]int{2, 3, 4, 5}, 3.5},
		test{[]int{5, 2, 8, 5}, 5},
		test{[]int{5, 10, 8, 5}, 6.5},
	}

	for _, v := range tests {
		got := CenteredAvg(v.data)
		if got != v.answer {
			t.Error("Error! Got: ", got, ", expected: ", v.answer)
		}
	}

}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{9000, 4, 10, 8, 6, 12, 12, 65, 234, 6, 34223, 4, 76354, 2334, 4, 63, 234, 6, 342, 14})
	}
}

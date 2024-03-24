package dog

import (
	"fmt"
	"testing"
)

// Testing function Years
func TestYears(t *testing.T) {
	got := Years(7)
	expected := 49

	if got != expected {
		t.Error("Expected: ", expected, " got: ", got)
	}
}

// Benchmarks run code multiple times to get avg run speed
// thus they need a loop for this testing
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

// Example for documentation
func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

// Testing function YearsTwo
func TestYearsTwo(t *testing.T) {
	got := YearsTwo(7)
	expected := 49

	if got != expected {
		t.Error("Expected: ", expected, " got: ", got)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

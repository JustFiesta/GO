package word

import (
	"fmt"
	"testing"
	"word/quote"
)

func TestCount(t *testing.T){
	got := Count("Hello world")
	expect := 2

	if got != expect {
		t.Error("Error! Got: ", got, " expected: ", expect)
	}
}

func ExampleCount() {
	fmt.Print(Count("Hello world"))
	// Output:
	// 2
}

func BenchmarkCount(b *testing.B){
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}


func BenchmarkUseCount(b *testing.B){
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
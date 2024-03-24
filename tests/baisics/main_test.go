package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(5, 7)
	want := 12
	if got != want {
		log.Fatal("error in addition, got: ", got, " expected: ", want)
		t.Error("Sum was incorrect, got: ", got, " expected: ", want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(5, 7)
	want := 35
	if got != want {
		log.Fatal("error in addition, got: ", got, " expected: ", want)
		t.Error("Sum was incorrect, got: ", got, " expected: ", want)
	}
}

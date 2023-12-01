package main

import "testing"

func TestPartOne(t *testing.T) {
	FileName = "partOneTestData.txt"
	got := PartOne()
	want := 142

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	FileName = "partTwoTestData.txt"
	got := PartTwo()
	want := 281

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo()
	}
}

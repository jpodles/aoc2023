package dayOne

import "testing"

func TestPartOne(t *testing.T) {
	FileName = "partOneTestData.txt"
	got := partOne()
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
	FileName = "day1.txt"
	for i := 0; i < b.N; i++ {
		partOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	FileName = "day1.txt"
	for i := 0; i < b.N; i++ {
		PartTwo()
	}
}

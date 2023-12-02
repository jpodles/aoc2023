package dayTwo

import "testing"

func TestPartOne(t *testing.T) {
	FileName = "d2p1td.txt"
	got := partOne()
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	FileName = "d2p1td.txt"
	got := partTwo()
	want := 2286

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	FileName = "day2.txt"
	for i := 0; i < b.N; i++ {
		partOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	FileName = "day2.txt"
	for i := 0; i < b.N; i++ {
		partTwo()
	}
}

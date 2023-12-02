package dayTwo

import "testing"

func TestPartOne(t *testing.T) {
	FileName = "d2p1td.txt"
	got := PartOne()
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

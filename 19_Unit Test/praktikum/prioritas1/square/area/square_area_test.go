package square

import "testing"

func TestSquareArea(t *testing.T) {

	if SquareArea(2) != "even square" {
		t.Error("Expected even square")
	}
	if SquareArea(3) != "odd square" {
		t.Error("Expected odd square")
	}
}

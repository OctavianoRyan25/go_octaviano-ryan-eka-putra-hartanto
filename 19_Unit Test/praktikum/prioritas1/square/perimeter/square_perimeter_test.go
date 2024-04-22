package square

import "testing"

func TestSquarePerimeter(t *testing.T) {

	if SquarePerimeter(2) != false {
		t.Error("Expected false")
	}
	if SquarePerimeter(11) != true {
		t.Error("Expected true")
	}
}

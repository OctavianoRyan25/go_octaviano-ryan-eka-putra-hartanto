package rectangle

import "testing"

func TestRectangleArea(t *testing.T) {

	if RectanglePerimeter(2, 2) != false {
		t.Error("Expected false")
	}
	if RectanglePerimeter(2, 3) != true {
		t.Error("Expected true")
	}
}

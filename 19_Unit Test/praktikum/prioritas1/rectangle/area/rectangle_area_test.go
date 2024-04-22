package rectangle

import "testing"

func TestRectangleArea(t *testing.T) {

	if RectangleArea(2, 3) != "even rectangle" {
		t.Error("Expected even rectangle")
	}
	if RectangleArea(3, 3) != "odd rectangle" {
		t.Error("Expected odd rectangle")
	}
}

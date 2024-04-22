package rectangle

func RectangleArea(length, width int) string {
	area := length * width

	if area%2 == 0 {
		return "even rectangle"
	} else {
		return "odd rectangle"
	}
}

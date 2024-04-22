package rectangle

func RectanglePerimeter(length, width int) bool {
	perimeter := 2 * (length + width)

	if perimeter%5 == 0 {
		return true
	} else {
		return false
	}
}

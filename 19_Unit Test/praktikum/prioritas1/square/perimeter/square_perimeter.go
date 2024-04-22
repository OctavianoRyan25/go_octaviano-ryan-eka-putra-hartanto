package square

func SquarePerimeter(s int) bool {
	perimeter := 4 * s

	if perimeter >= 40 {
		return true
	} else {
		return false
	}
}

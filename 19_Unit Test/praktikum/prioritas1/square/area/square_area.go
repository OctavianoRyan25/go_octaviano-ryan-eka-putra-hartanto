package square

func SquareArea(s int) string {
	area := s * s

	if area%2 == 0 {
		return "even square"
	} else {
		return "odd square"
	}
}

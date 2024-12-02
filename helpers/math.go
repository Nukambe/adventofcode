package helpers

func AbsoluteInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

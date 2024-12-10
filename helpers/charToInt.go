package helpers

import "strconv"

func CharToInt(c uint8) int {
	i, _ := strconv.Atoi(string(c))
	return i
}

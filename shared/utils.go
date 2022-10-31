package shared

import "strconv"

func StringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		number = 0
	}
	return number
}

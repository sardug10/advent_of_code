package utils

import "strconv"

func Contains(arr []string, ele int) bool {
	for _, char := range arr {
		currentNum, _ := strconv.Atoi(char)
		if currentNum == ele {
			return true
		}
	}
	return false
}

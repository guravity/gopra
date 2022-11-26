package strcount

import "strings"

func StrCount(str string, char string) int {
	count := strings.Count(str, char)
	return count
}
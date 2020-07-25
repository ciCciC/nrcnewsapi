package util

import "strings"

func IsEmpty(value string) bool {
	return value == ""
}

func TrimText(text string) string {
	trimmedSpaces := strings.TrimSpace(text)
	trimmedLeft := strings.TrimLeft(trimmedSpaces, " ")
	trimmedFinal := strings.TrimRight(trimmedLeft, " ")
	return trimmedFinal
}

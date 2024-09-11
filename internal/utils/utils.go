package utils

import "github.com/google/uuid"

func GetUniqueId() string {
	id := uuid.New()
	return GetSubString(id.String(), 0, 5)
}

func GetLeftString(input string, length int) string {
	return GetSubString(input, 0, length)
}

func GetSubString(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if length >= len(asRunes) {
		length = len(asRunes)
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

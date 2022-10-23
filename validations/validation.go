package validations

import "unicode"

func IsValidRotorPosition(k rune) bool {
	return unicode.IsUpper(k)
}

func IsValidPlugBoardPin(k rune) bool {
	return unicode.IsUpper(k)
}

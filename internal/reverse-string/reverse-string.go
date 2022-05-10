package reverse_string

import (
	"unicode"
)

// Main function
func reverseString(s string) string {
	textRunes := []rune(s)
	textRunesLength := len(textRunes)
	if textRunesLength <= 1 {
		return s
	}

	i, j := 0, 0
	for i < textRunesLength && j < textRunesLength {
		j = i + 1
		for j < textRunesLength && isMark(textRunes[j]) {
			j++
		}

		if isMark(textRunes[j-1]) {
			// Reverses Combined Characters
			reverse(textRunes[i:j], j-i)
		}

		i = j
	}

	// Reverses the entire array
	reverse(textRunes, textRunesLength)

	return string(textRunes)
}

// Reverse helping function
func reverse(runes []rune, length int) {
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// isMark determines whether the rune is a marker
func isMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}

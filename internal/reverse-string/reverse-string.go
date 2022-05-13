package reverse_string

func reverseString(s string) string {
	// string to rune slice
	runes := []rune(s)

	var result []rune

	// iterate rune slice in reverse, build result
	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}

	return string(result)
}

package iteration

import "strings"

// Repeat builds a string by repeating a character a specified number of times
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

package iteration

import "strings"

// Repeat builds a string of a repeated character a given number of times
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

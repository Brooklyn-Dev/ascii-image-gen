package utils

import (
	"strings"
)

// Reverses the characters of a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Creates a square string
func CreateSquare(size int, char string, aspectRatio float64) string {
	width := int(float64(size) / aspectRatio)
	
    line := strings.Repeat(char, width)
    square := strings.Repeat(line + "\n", size)

	return square
}
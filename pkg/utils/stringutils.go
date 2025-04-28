package utils

import (
	"path/filepath"
	"regexp"
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

// Creates a filename string
func CreateSaveFilename(imgPath string, newExt string) string {
	filename := filepath.Base(imgPath) 
    ext := filepath.Ext(imgPath)
    return strings.TrimSuffix(filename, ext) + "-ascii-art" + newExt
}

var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)
// Removes all ANSI codes from a string
func StripANSI(str string) string {
    return ansiRegexp.ReplaceAllString(str, "")
}
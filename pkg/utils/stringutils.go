package utils

import (
	"fmt"
	imageColor "image/color"
	"path/filepath"
	"strconv"
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

func StringToRGBA(rgbaStr string) (imageColor.RGBA, error) {
	strValues := strings.Split(rgbaStr, ",")
	if len(strValues) != 4 {
		return imageColor.RGBA{}, fmt.Errorf("string must have four 0-255 integer values separated by commas")
	}

	var rgba [4]uint8

	for i, strVal := range strValues {
		intVal, err := strconv.ParseUint(strVal, 10, 8)
		if err != nil {
			return imageColor.RGBA{}, fmt.Errorf("invalid value %q must be a integer between 0-255", strVal)
		}

		rgba[i] = uint8(intVal)
	}

	return imageColor.RGBA{rgba[0], rgba[1], rgba[2], rgba[3]}, nil
}
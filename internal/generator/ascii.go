package generator

import (
	"image"
	"strings"
)

// Config for generator options
type Config struct {
	Width int
}

const charRamp = "@%#*+=-:."

// Generates ASCII string from image
func ImageToASCII(img image.Image, config Config) string {
	bounds := img.Bounds()
	var builder strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			a8 := uint8(a >> 8)
			if a8 < 16 {
				builder.WriteString(" ")
				continue
			}

			// Convert to 8-bit
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)

			// Weighted greyscale
			grey := 0.299 * float64(r8) + 0.587 * float64(g8) + 0.144 * float64(b8)

			// Map to char index
			index := int(grey * float64(len(charRamp) - 1) / 255)
			if index >= len(charRamp) {
				index = len(charRamp) - 1
			} else if index < 0 {
				index = 0
			}

			char := string(charRamp[index])
			builder.WriteString(char)
		}

		builder.WriteString("\n")
	}

	return builder.String()
}
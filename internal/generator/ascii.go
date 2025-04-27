package generator

import (
	"image"
	"strings"

	"github.com/gookit/color"

	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
)

// Config for generator options
type Config struct {
	Colour bool
	Greyscale bool
	Invert bool
	Width int
}

const charRamp = ".:-=+*#%@"

// Generates ASCII string from image
func ImageToASCII(img image.Image, config Config) string {
	bounds := img.Bounds()
	var builder strings.Builder

	brightnessRamp := charRamp
	if config.Invert {
		brightnessRamp = utils.ReverseString(charRamp)
	}

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

			grey, idx := computeBrightness(r8, g8, b8)
			grey8 := uint8(grey)

			char := string(brightnessRamp[idx])

			if config.Colour {
				builder.WriteString(color.RGB(r8, g8, b8).Sprint((char)))
			} else if config.Greyscale {
				builder.WriteString(color.RGB(grey8, grey8, grey8).Sprint((char)))
			} else {
				builder.WriteString(char)
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

// Calculates the greyscale value and maps it to a character index
func computeBrightness(r, g, b uint8) (float64, int) {
	// Weighted greyscale
	grey := 0.299 * float64(r) + 0.587 * float64(g) + 0.114 * float64(b)

	// Map to char index
	index := int(grey * float64(len(charRamp) - 1) / 255)
	if index >= len(charRamp) {
		index = len(charRamp) - 1
	} else if index < 0 {
		index = 0
	}

	return grey, index
}
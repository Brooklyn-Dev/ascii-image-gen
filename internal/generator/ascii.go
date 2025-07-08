package generator

import (
	"image"
	imageColor "image/color"
	"image/draw"
	"strings"

	"github.com/gookit/color"
	"github.com/leaanthony/go-ansi-parser"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
)

// Config for generator options
type Config struct {
	AspectRatio float64
	Colour bool
	Complex bool
	FlipX bool
	FlipY bool
	Greyscale bool
	Invert bool
	Negative bool
	SaveBG string
	SaveDir string
	SavePNG bool
	SaveText bool
	Width int
}

// Ramps taken from http://paulbourke.net/dataformats/asciiart/
const simpleRamp = ".:-=+*#%@"
const complexRamp = ".'`^\",:;Il!i><~+_-?]}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

// Generates ASCII string from image
func ImageToASCII(img image.Image, config Config) string {
	bounds := img.Bounds()
	var builder strings.Builder

	brightnessRamp := simpleRamp
	if config.Complex {
		brightnessRamp = complexRamp	
	}
	if config.Invert {
		brightnessRamp = utils.ReverseString(brightnessRamp)
	}

	// Iterate through image pixels indicies
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Apply flips
			imgX := x
			if config.FlipX {
				imgX = bounds.Max.X - 1 - (x - bounds.Min.X)
			} 

			imgY := y
			if config.FlipY {
				imgY = bounds.Max.Y - 1 - (y - bounds.Min.Y)
			} 

			// Get RGBA value
			r, g, b, a := img.At(imgX, imgY).RGBA()

			// Ignore transparent (or almost transparent) pixels
			a8 := uint8(a >> 8)
			if a8 < 16 {
				builder.WriteString(" ")
				continue
			}

			// Convert RBG to 8-bit
			r8, g8, b8 := uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)

			// Negate if true
			if config.Negative {
				r8, g8, b8 = ^r8, ^g8, ^b8
			}

			// Calculate brightness and map to char index
			grey, idx := computeBrightness(r8, g8, b8, brightnessRamp)
			grey8 := uint8(grey)

			char := string(brightnessRamp[idx])

			// Write the character in the appropriate colour
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
func computeBrightness(r, g, b uint8, charRamp string) (float64, int) {
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

// Generates ASCII image from ASCII string
func ASCIIToImageArt(ascii string, config Config) (*image.RGBA, error) {
	lines := strings.Split(ascii, "\n")
	fontFace := basicfont.Face7x13
	charWidth := fontFace.Advance
	charHeight := fontFace.Height

	imgWidth := charWidth * config.Width
	imgHeight := charHeight * len(lines)
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	bgCol, err := utils.StringToRGBA(config.SaveBG)
	if err != nil {
		return nil, err
	}
	imgBg := image.NewUniform(bgCol)

	draw.Draw(img, img.Bounds(), imgBg, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  img,
		Face: fontFace,
	}

	// Plain ASCII
	if !(config.Colour || config.Greyscale) {
		d.Src = image.NewUniform(imageColor.White)

		for y, line := range lines {
			d.Dot = fixed.P(0, (y + 1) * charHeight)
			d.DrawString(line)
		}

		return img, nil
	}

	// Coloured or Greyscale ASCII
	text, err := ansi.Parse(ascii)
	if err != nil {
		return nil, err
	}

	x, y := 0, 0
	for _, seg := range text {
		// Get colour
		var col imageColor.Color = imageColor.White
		if seg.FgCol != nil {
			rgb := seg.FgCol.Rgb
			col = imageColor.RGBA{rgb.R, rgb.G, rgb.B, 255}
		}
	
		// Set colour
		d.Src = image.NewUniform(col)
	
		// Draw each character
		lines := strings.Split(seg.Label, "\n")
		for i, line := range lines {
			// Reset X on new lines
			if i > 0 {
				x = 0
				y++
			}

			d.Dot = fixed.P(x * charWidth, (y + 1) * charHeight)
			d.DrawString(line)

			x += len(line)
		}
	}

	return img, nil
}
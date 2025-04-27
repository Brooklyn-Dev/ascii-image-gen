package generator

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

const aspectCorrection = 0.4

// Loads an image from a file path
func loadImage(imgPath string) (image.Image, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if strings.HasSuffix(strings.ToLower(imgPath), ".webp") {
		img, err := webp.Decode(file)
		if err != nil {
			return nil, err
		}
		return img, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, err
}

// Resize an image based on config
func resizeImage(img image.Image, config Config) image.Image {
	origBounds := img.Bounds()
	origWidth, origHeight := origBounds.Dx(), origBounds.Dy()

	aspectRatio := float64(origHeight) / float64(origWidth)

	newWidth := config.Width
	newHeight := int(aspectRatio * float64(newWidth) * aspectCorrection)

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), img, origBounds, draw.Over, nil)

	return newImg
}

// Performs the entire process to generate ASCII string from file path
func ImageFileToASCII(imgPath string, config Config) (string, error) {
	// Load image
	img, err := loadImage(imgPath)
	if err != nil {
		return "", err
	}
	
	// Resize image
	img = resizeImage(img, config)
	
	// Convert to ASCII
	return ImageToASCII(img, config), nil
}
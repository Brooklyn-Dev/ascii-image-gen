package cli

import (
	"flag"
	"fmt"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
)

// Parses command line flags and returns config
func ParseFlags() (*generator.Config, error) {
	config := generator.Config{
		Colour: false,
		Complex: false,
		Greyscale: false,
		Invert: false,
		Negative: false,
		Width: 100,
	}

	// Define flags
	flag.BoolVar(&config.Colour, "colour", false, "Colour the generated ASCII")
	flag.BoolVar(&config.Complex, "complex", false, "Use a more detailed character ramp")
	flag.BoolVar(&config.FlipX, "flip-x", false, "Horizontally flip the generated ASCII")
	flag.BoolVar(&config.FlipY, "flip-y", false, "Vertically flip the generated ASCII")
	flag.BoolVar(&config.Greyscale, "greyscale", false, "Grey the generated ASCII")
	flag.BoolVar(&config.Invert, "invert", false, "Invert the character ramp")
	flag.BoolVar(&config.Negative, "negative", false, "Negate colours of all characters")
	flag.IntVar(&config.Width, "width", 100, "Width of the generated ASCII")
	
	// Parse flags
	flag.Parse()

	if config.Colour && !utils.SupportsColour() {
        return nil, fmt.Errorf("terminal does not support colours")
	}

	if config.Colour && config.Greyscale {
        return nil, fmt.Errorf("cannot use -colour and -greyscale together")
    }

	if config.Width <= 0 {
        return nil, fmt.Errorf("width must be greater than 0")
    }

	return &config, nil
}
package cli

import (
	"flag"
	"fmt"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
)

// Parses command line flags and returns config
func ParseFlags() (*generator.Config, error) {
	config := generator.Config{
		Colour: false,
		Greyscale: false,
		Invert: false,
		Width: 100,
	}

	// Define flags
	flag.BoolVar(&config.Colour, "colour", false, "Colour the generated ascii")
	flag.BoolVar(&config.Greyscale, "greyscale", false, "Grey the generated ascii")
	flag.BoolVar(&config.Invert, "invert", false, "Invert the character ramp")
	flag.IntVar(&config.Width, "width", 100, "Width of the generated ascii")
	
	// Parse flags
	flag.Parse()

	if config.Colour && config.Greyscale {
        return nil, fmt.Errorf("cannot use -colour and -greyscale together")
    }

	if config.Width <= 0 {
        return nil, fmt.Errorf("width must be greater than 0")
    }

	return &config, nil
}
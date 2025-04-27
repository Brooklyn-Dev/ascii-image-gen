package cli

import (
	"flag"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
)

// Parses command line flags and returns config
func ParseFlags() (generator.Config) {
	config := generator.Config{
		Width: 100,
	}

	// Define flags
	flag.IntVar(&config.Width, "width", 100, "Width of the generated ascii")
	
	// Parse flags
	flag.Parse()

	return config
}
package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
)

// Parses command line flags and returns config
func ParseFlags() (*generator.Config, error) {
	config := generator.Config{
		AspectRatio: 0.5,
		Colour: false,
		Complex: false,
		FlipX: false,
		FlipY: false,
		Greyscale: false,
		Invert: false,
		Negative: false,
		SaveBG: "0,0,0,0",
		SaveDir: "",
		SavePNG: false,
		SaveText: false,
		Verbose: false,
		Version: false,
		Width: 100,
	}

	// Define short flags
	flag.Float64Var(&config.AspectRatio, "a", 0.5, "Character aspect ratio (width:height) for your terminal font")
	flag.BoolVar(&config.Colour, "c", false, "Colour the generated ASCII")
	flag.BoolVar(&config.Complex, "C", false, "Use a more detailed character ramp")
	flag.BoolVar(&config.FlipX, "x", false, "Horizontally flip the generated ASCII")
	flag.BoolVar(&config.FlipY, "y", false, "Vertically flip the generated ASCII")
	flag.BoolVar(&config.Greyscale, "g", false, "Grey the generated ASCII")
	flag.BoolVar(&config.Invert, "i", false, "Invert the character ramp")
	flag.BoolVar(&config.Negative, "n", false, "Negate colours of all characters")
	flag.StringVar(&config.SaveDir, "d", "", "Save directory of saved files")
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose logging")
	flag.BoolVar(&config.Version, "V", false, "Shows version")
	flag.IntVar(&config.Width, "w", 100, "Width of the generated ASCII")

	// Define long flags
	flag.Float64Var(&config.AspectRatio, "aspect-ratio", 0.5, "Character aspect ratio (width:height) for your terminal font")
	flag.BoolVar(&config.Colour, "color", false, "Color the generated ASCII (alias)")
	flag.BoolVar(&config.Colour, "colour", false, "Colour the generated ASCII")
	flag.BoolVar(&config.Complex, "complex", false, "Use a more detailed character ramp")
	flag.BoolVar(&config.FlipX, "flip-x", false, "Horizontally flip the generated ASCII")
	flag.BoolVar(&config.FlipY, "flip-y", false, "Vertically flip the generated ASCII")
	flag.BoolVar(&config.Greyscale, "grayscale", false, "Gray the generated ASCII (alias)")
	flag.BoolVar(&config.Greyscale, "greyscale", false, "Grey the generated ASCII")
	flag.BoolVar(&config.Invert, "invert", false, "Invert the character ramp")
	flag.BoolVar(&config.Negative, "negative", false, "Negate colours of all characters")
	flag.StringVar(&config.SaveBG, "save-bg", "0,0,0,0", "Set background RGBA for saved PNG files")
	flag.StringVar(&config.SaveDir, "save-dir", "", "Save directory of saved files")
	flag.BoolVar(&config.SavePNG, "save-png", false, "Save generated ASCII in png file(s)")
	flag.BoolVar(&config.SaveText, "save-text", false, "Save generated ASCII in text file(s)")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&config.Version, "version", false, "Shows version")
	flag.IntVar(&config.Width, "width", 100, "Width of the generated ASCII")

	calibratePtr := flag.Bool("calibrate", false, "Calibrate to help manually determine aspect ratio")
	calibrationSizePtr := flag.Int("calibration-size", 16, "Size of the calibration test")

	// Parse flags
	flag.Parse()

	if config.Version {
		println("ascii-image-gen version 1.0.0")
		os.Exit(0)
	}

	if *calibratePtr {
		// Calibration size range validation
		if *calibrationSizePtr <= 0 {
			return nil, fmt.Errorf("calibration-size must be greater than 0")
		}
		if *calibrationSizePtr >= 33 {
			return nil, fmt.Errorf("calibration-size must be less than 33")
		}

		// Unicode support check
		testChar := "█"
		if !utils.SupportsUnicode(testChar) {
			testChar = "@"
		}

		// Draw calibration square
		square := utils.CreateSquare(*calibrationSizePtr, testChar, config.AspectRatio)
		fmt.Print(square)

		os.Exit(0)
	}

	_, err := utils.StringToRGBA(config.SaveBG)
	if err != nil {
		return nil, fmt.Errorf("invalid -save-bg value: %v", err)
	}

	// Flag combination validations
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
package main

import (
	"fmt"
	"os"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/cli"
	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
)

func main() {
	// Parse command line flags and args
	config, err := cli.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgPaths, err := cli.ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Process each image
	for _, imgPath := range imgPaths {
		fmt.Printf("Processing: %s\n", imgPath)

		// Generate ASCII
		ascii, err := generator.ImageFileToASCII(imgPath, *config)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", imgPath, err)
			continue
		}

		// Print result
		fmt.Println(ascii)
	}
}
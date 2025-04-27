package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/cli"
	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
)

func main() {
	// Remove datetime from logs
	log.SetFlags(0)

	// Parse command line flags and args
	config, err := cli.ParseFlags()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	imgPaths, err := cli.ParseArgs()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Process each image
	for _, imgPath := range imgPaths {
		log.Printf("Processing: %s\n", imgPath)

		// Generate ASCII
		ascii, err := generator.ImageFileToASCII(imgPath, *config)
		if err != nil {
			log.Printf("Error processing %s: %v\n", imgPath, err)
			continue
		}

		// Print result
		fmt.Println(ascii)
	}
}
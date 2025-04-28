package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/cli"
	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
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
		
		// Save if applicable
		if config.SaveText {
			filename := utils.CreateSaveFilename(imgPath, ".txt")
			
			if config.Colour {
				ascii = utils.StripANSI(ascii)
			}
			
			saveDir := ""
			if config.SaveDir != "" {
				if !utils.IsValidPath(config.SaveDir) {
					log.Println(fmt.Errorf("invalid output directory: %s", config.SaveDir))
					os.Exit(1)
				}

				saveDir = config.SaveDir
			}
			
			savePath := filepath.Join(saveDir, filename)
			log.Printf("Saving: %s\n", savePath)

			err := utils.SaveAsText(ascii, savePath)
			if err != nil {
				log.Printf("Error saving %s: %v\n", savePath, err)
				continue
			}	
		}
	}
}
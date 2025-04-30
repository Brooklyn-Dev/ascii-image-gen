package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Brooklyn-Dev/ascii-image-gen/internal/cli"
	"github.com/Brooklyn-Dev/ascii-image-gen/internal/generator"
	"github.com/Brooklyn-Dev/ascii-image-gen/pkg/utils"
	"github.com/leaanthony/go-ansi-parser"
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
		if config.SavePNG || config.SaveText {
			// Get save directory
			saveDir := ""
			if config.SaveDir != "" {
				if !utils.IsValidPath(config.SaveDir) {
					log.Println(fmt.Errorf("invalid output directory: %s", config.SaveDir))
					os.Exit(1)
				}

				saveDir = config.SaveDir
			}

			if config.SaveText {
				textAscii := ascii
				if config.Colour {
					log.Println("Cleansing text save ASCII")
					textAscii, err = ansi.Cleanse(ascii)
					if err != nil {
						log.Printf("Error cleansing text save ASCII: %v\n", err)
						continue					
					}
				}
	
				log.Println("Preparing text save path")
				savePath := filepath.Join(saveDir, utils.CreateSaveFilename(imgPath, ".txt"))
				savePath, err := utils.FindAvaliablePath(savePath)
				if err != nil {
					log.Printf("Error preparing text save path: %v\n", err)
					continue
				}

				log.Printf("Saving text: %s\n", savePath)
				err = utils.SaveAsText(textAscii, savePath)
				if err != nil {
					log.Printf("Error saving text %s: %v\n", savePath, err)
					continue
				}	
			}

			if config.SavePNG {
				log.Println("Preparing PNG save path")
				savePath := filepath.Join(saveDir, utils.CreateSaveFilename(imgPath, ".png"))
				savePath, err := utils.FindAvaliablePath(savePath)
				if err != nil {
					log.Printf("Error preparing PNG save path: %v\n", err)
					continue
				}

				log.Printf("Generating PNG: %s\n", savePath)
				img, err := generator.ASCIIToImageArt(ascii, *config)
				if err != nil {
					log.Printf("Error generating PNG %s: %v\n", savePath, err)
					continue
				}
				
				log.Printf("Saving PNG: %s\n", savePath)
				err = utils.SaveAsPNG(img, savePath)
				if err != nil {
					log.Printf("Error saving PNG %s: %v\n", savePath, err)
					continue
				}	
			}
		}
	}
}
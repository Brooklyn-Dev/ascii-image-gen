package cli

import (
	"flag"
	"fmt"
)

// Parses command line args and returns image paths
func ParseArgs() ([]string, error) {
	// Parse args
	imgPaths := flag.Args()

	if len(imgPaths) == 0 {
		return nil, fmt.Errorf("usage: ascii-image-gen [options] image1.jpg [image2.jpg ...]")
	}

	return imgPaths, nil
}
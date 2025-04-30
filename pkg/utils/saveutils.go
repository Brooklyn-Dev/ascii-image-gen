package utils

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"image/png"
)

// Check if specified path is valid
func IsValidPath(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

// Auto-increments filename for save path conflicts
func FindAvaliablePath(basePath string) (string, error) {
    dir := filepath.Dir(basePath)
    ext := filepath.Ext(basePath)
    name := filepath.Base(basePath[:len(basePath) - len(ext)])

    newPath := basePath
    i := 1
    for {
        _, err := os.Stat(newPath)

        if os.IsNotExist(err) {
            return newPath, nil
        }

        if err != nil {
            return "", err
        }

        newPath = filepath.Join(dir, fmt.Sprintf("%s-%d%s", name, i, ext))
        i++
    }
}

// Save string content in a TXT file
func SaveAsText(content string, path string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

// Save image in a PNG file
func SaveAsPNG(img image.Image, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    return png.Encode(file, img)
}
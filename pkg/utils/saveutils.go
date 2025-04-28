package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

// Save string content in a text file
func SaveAsText(content string, path string) error {
    return os.WriteFile(path, []byte(content), 0644)
}
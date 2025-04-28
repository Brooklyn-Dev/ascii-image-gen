package utils

import "os"

// Check if specified path is valid
func IsValidPath(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

// Save string content in a text file
func SaveAsText(content string, path string) error {
    return os.WriteFile(path, []byte(content), 0644)
}
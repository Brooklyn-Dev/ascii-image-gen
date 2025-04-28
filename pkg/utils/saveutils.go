package utils

import "os"

// Save string content in a text file
func SaveAsText(content string, path string) error {
    return os.WriteFile(path, []byte(content), 0644)
}
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Starting...")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".gif") {
			// If the file is a .gif, move it to the current directory
			newPath := filepath.Join(".", info.Name())
			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				err = os.Rename(path, newPath)
				if err != nil {
					fmt.Printf("Failed to move file: %s\n", err)
				} else {
					fmt.Printf("Moved file: %s to %s\n", path, newPath)
				}
			} else {
				// Iterate number until we find a file that doesn't exist
				i := 1
				for {
					newPath = filepath.Join(".", fmt.Sprintf("%s-%d.gif", strings.TrimSuffix(info.Name(), ".gif"), i))
					if _, err := os.Stat(newPath); os.IsNotExist(err) {
						break
					}
					i++
				}

				err = os.Rename(path, newPath)
				if err != nil {
					fmt.Printf("Failed to move file: %s\n", err)
				} else {
					fmt.Printf("Moved file: %s to %s\n", path, newPath)
				}

				fmt.Printf("File exists, renamed to %s\n", newPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return
	}
}

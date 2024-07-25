package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "./content/posts"

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		newFileName := strings.ToLower(strings.ReplaceAll(info.Name(), " ", "-"))
		newPath := filepath.Join(filepath.Dir(path), newFileName)

		if newPath != path {
			fmt.Printf("Renaming %s to %s\n", path, newPath)
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Println("Error renaming file:", err)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

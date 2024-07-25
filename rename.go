package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rename <directory>")
		os.Exit(1)
	}

	dir := os.Args[1]

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

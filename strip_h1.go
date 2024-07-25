package main

import (
	"log"
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
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			processFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", dir, err)
	}
}

func processFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", path, err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "# ") {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	err = os.WriteFile(path, []byte(newContent), 0o644)
	if err != nil {
		log.Printf("Error writing file %s: %v\n", path, err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type FrontMatter struct {
	Title   string `yaml:"title"`
	LastMod string `yaml:"lastmod"`
}

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

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		log.Printf("No YAML front matter found in %s\n", path)
		return
	}

	var fm FrontMatter
	err = yaml.Unmarshal([]byte(parts[1]), &fm)
	if err != nil {
		log.Printf("Error parsing YAML front matter in %s: %v\n", path, err)
		return
	}

	newFrontMatter, err := yaml.Marshal(&fm)
	if err != nil {
		log.Printf("Error marshaling new front matter for %s: %v\n", path, err)
		return
	}

	newContent := fmt.Sprintf("---\n%s---\n%s", newFrontMatter, parts[2])
	err = os.WriteFile(path, []byte(newContent), 0o644)
	if err != nil {
		log.Printf("Error writing file %s: %v\n", path, err)
	}
}

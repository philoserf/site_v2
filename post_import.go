package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// FrontMatter struct from strip_frontmatter.go
type FrontMatter struct {
	Title   string `yaml:"title"`
	LastMod string `yaml:"lastmod"`
	Date    string `yaml:"date"`
}

// Add other type definitions from rename_imports.go, strip_footer.go here

func main() {
	path := "./content/posts"

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			processFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", path, err)
	}
}

func processFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file %q: %v\n", path, err)
	}

	// Call functions to process the file content
	content = stripFrontMatter(content)
	content = stripH1(content)
	content = stripFooter(content)

	err = os.WriteFile(path, content, 0o644)
	if err != nil {
		log.Fatalf("Error writing file %q: %v\n", path, err)
	}

	err = renameImports(path)
	if err != nil {
		log.Fatalf("Error renaming file %q: %v\n", path, err)
	}
}

// Function from strip_frontmatter.go
func stripFrontMatter(content []byte) []byte {
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		log.Printf("No YAML front matter found\n")
		return content
	}

	var fm FrontMatter
	err := yaml.Unmarshal([]byte(parts[1]), &fm)
	if err != nil {
		log.Printf("Error parsing YAML front matter: %v\n", err)
		return content
	}

	newFrontMatter, err := yaml.Marshal(&fm)
	if err != nil {
		log.Printf("Error marshaling new front matter: %v\n", err)
		return content
	}

	newContent := fmt.Sprintf("---\n%s---\n%s", newFrontMatter, parts[2])
	return []byte(newContent)
}

// Function from strip_h1.go
func stripH1(content []byte) []byte {
	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "# ") {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	return []byte(newContent)
}

// Function from strip_footer.go
func stripFooter(content []byte) []byte {
	lines := strings.Split(string(content), "\n")
	var newLines []string
	found := false

	for _, line := range lines {
		if strings.TrimSpace(line) == "***" {
			found = true
			break
		}
		newLines = append(newLines, line)
	}

	if !found {
		log.Printf("No '***' line found\n")
		return content
	}

	newContent := strings.Join(newLines, "\n")
	return []byte(newContent)
}

// Function from rename_imports.go
func renameImports(dir string) error {
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
	return nil
}

// Add other functions from rename_imports.go, strip_footer.go here

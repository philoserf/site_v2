package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// FrontMatter represents the YAML front matter of an article.
type FrontMatter struct {
	Title   string `yaml:"title"`
	LastMod string `yaml:"lastmod"`
	Date    string `yaml:"date"`
}

func main() {
	const path = "./content/posts"

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			if err := processFile(path); err != nil {
				return fmt.Errorf("error processing file %q: %w", path, err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("error walking the path %q: %v", path, err)
	}
}

// processFile coordinates the processing of a file.
func processFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", path, err)
	}

	content = stripFrontMatter(content)
	content = stripH1(content)
	content = stripFooter(content)

	const ownerReadWrite = 0o600

	err = os.WriteFile(path, content, ownerReadWrite)
	if err != nil {
		return fmt.Errorf("error writing file %q: %w", path, err)
	}

	err = renameFile(path)
	if err != nil {
		return fmt.Errorf("error renaming file %q: %w", path, err)
	}

	return nil
}

// stripFrontMatter removes the YAML front matter from the content.
func stripFrontMatter(content []byte) []byte {
	const (
		frontMatterParts = 3
		delimiter        = "---"
	)

	parts := strings.SplitN(string(content), delimiter, frontMatterParts)
	if len(parts) < frontMatterParts {
		return content
	}

	var frontMatter FrontMatter

	err := yaml.Unmarshal([]byte(parts[1]), &frontMatter)
	if err != nil {
		return content
	}

	newFrontMatter, err := yaml.Marshal(&frontMatter)
	if err != nil {
		return content
	}

	newContent := fmt.Sprintf("%s\n%s%s\n%s", delimiter, newFrontMatter, delimiter, parts[2])

	return []byte(newContent)
}

// stripH1 removes the first H1 header from the content.
func stripH1(content []byte) []byte {
	const headerPrefix = "# "

	lines := strings.Split(string(content), "\n")

	newLines := make([]string, 0, len(lines))

	for _, line := range lines {
		if !strings.HasPrefix(line, headerPrefix) {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")

	return []byte(newContent)
}

// stripFooter removes the markdown footer from the content.
func stripFooter(content []byte) []byte {
	const footerMarker = "***"

	lines := strings.Split(string(content), "\n")

	newLines := make([]string, 0, len(lines))

	found := false

	for _, line := range lines {
		if strings.TrimSpace(line) == footerMarker {
			found = true

			break
		}

		newLines = append(newLines, line)
	}

	if !found {
		return content
	}

	newContent := strings.Join(newLines, "\n")

	return []byte(newContent)
}

// renameFile renames the file to lowercase and replaces spaces with hyphens.
func renameFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%w: error getting file info for %q", err, path)
	}

	newFileName := strings.ToLower(strings.ReplaceAll(info.Name(), " ", "-"))
	newPath := filepath.Join(filepath.Dir(path), newFileName)

	if newPath != path {
		err := os.Rename(path, newPath)
		if err != nil {
			return fmt.Errorf("error renaming file %q to %q: %w", path, newPath, err)
		}
	}

	return nil
}

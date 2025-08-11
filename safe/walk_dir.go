package safe

import (
	"os"
	"path/filepath"
	"strings"
)

// WalkDir recursively walks a directory, skipping entries based on .hideignore patterns.
func WalkDir(root string) ([]string, error) {
	patterns := loadHideIgnore(root)
	var results []string

	var walk func(string) error
	walk = func(path string) error {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			name := entry.Name()
			fullPath := filepath.Join(path, name)

			if shouldSkip(name, patterns) {
				continue
			}

			results = append(results, fullPath)

			if entry.IsDir() {
				if err := walk(fullPath); err != nil {
					return err
				}
			}
		}
		return nil
	}

	err := walk(root)
	return results, err
}

// loadHideIgnore reads .hideignore patterns from root directory.
func loadHideIgnore(root string) []string {
	filePath := filepath.Join(root, ".hideignore")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil // no .hideignore
	}
	lines := strings.Split(string(data), "\n")
	var patterns []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}
	return patterns
}

// shouldSkip checks if a name matches any ignore pattern.
func shouldSkip(name string, patterns []string) bool {
	if strings.HasPrefix(name, ".") ||
		strings.HasSuffix(name, ".zip") ||
		strings.HasSuffix(name, ".enc") {
		return true
	}
	for _, pattern := range patterns {
		match, _ := filepath.Match(pattern, name)
		if match {
			return true
		}
	}
	return false
}

package safe

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func deriveKey(password string) *[32]byte {
	hash := sha256.Sum256([]byte(password))
	var key [32]byte
	copy(key[:], hash[:])
	return &key
}

type Options struct {
	Mode     string
	Src      string
	Dest     string
	Password string
}

func isHidden(path string) bool {
	base := filepath.Base(path)
	return strings.HasPrefix(base, ".") && len(base) > 1
}

func ReadHideRules() []string {
	lines := []string{}
	filePath := ".hideignore"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("No \".hide\" file [working on all files]", err)
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines
}

func GetFilteredPaths(startDir string) ([]string, []string, error) {
	hide_rules := ReadHideRules()
	files, err := WalkDir(startDir)
	src_files := []string{}
	dest_files := []string{}

	fmt.Println("Hide rules:", hide_rules)

	for _, file_path := range files {
		dest_path := fmt.Sprintf("%s.enc", file_path)
		src_files = append(src_files, file_path)
		dest_files = append(dest_files, dest_path)
	}
	return src_files, dest_files, err
}

func Consume(args ...any) {

}

func GetEncryptedFiles(startDir string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip hidden files/folders
		if isHidden(path) {
			if info.IsDir() {
				return filepath.SkipDir // skip entire hidden directory
			}
			return nil // skip hidden file
		}
		// Get relative path
		rel, err := filepath.Rel(startDir, path)
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(rel) == ".enc" {
			files = append(files, rel)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Reading files failed")
	}

	return files, err
}

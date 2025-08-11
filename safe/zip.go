package safe

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func ZipDir(opts Options) error {
	src_files, _, err := GetFilteredPaths(opts.Src)
	if err != nil {
		fmt.Println("Can not read some files", err)
		return nil
	}

	zip_file_path := "out.zip"
	zip_file, err := os.Create(zip_file_path)
	if err != nil {
		return fmt.Errorf("cannot create zip: %w", err)
	}
	defer zip_file.Close()

	zipWriter := zip.NewWriter(zip_file)
	defer zipWriter.Close()

	for _, relPath := range src_files {
		fullPath := filepath.Join(opts.Src, relPath)

		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			return fmt.Errorf("stat error for %s: %w", relPath, err)
		}

		zipHeader, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return fmt.Errorf("header error for %s: %w", relPath, err)
		}

		// Store file with relative path inside zip
		zipHeader.Name = relPath
		zipHeader.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(zipHeader)
		if err != nil {
			return fmt.Errorf("zip create error for %s: %w", relPath, err)
		}

		if !fileInfo.IsDir() {
			srcFile, err := os.Open(fullPath)
			if err != nil {
				return fmt.Errorf("open error for %s: %w", relPath, err)
			}
			_, err = io.Copy(writer, srcFile)
			srcFile.Close()
			if err != nil {
				return fmt.Errorf("copy error for %s: %w", relPath, err)
			}
		}
	}

	zipWriter.Close()
	time.Sleep(time.Duration(time.Duration.Seconds(2)))

	// Encrypt the zip
	zip_out_path := fmt.Sprintf("%s.enc", zip_file_path)
	err = EncryptFile(zip_file_path, zip_out_path, opts.Password)
	if err != nil {
		fmt.Printf("Encryption failed for %s: %v\n", opts.Src, err)
	}

	time.Sleep(time.Duration(time.Duration.Seconds(2)))
	os.Remove(zip_file_path)

	fmt.Printf("Zipped and encrypted %d files\n", len(src_files))
	return nil
}

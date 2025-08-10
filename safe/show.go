package safe

import (
	"fmt"
	"os"
	"path/filepath"
)

func ShowDir(opts Options) error {
	src_files, err := GetEncryptedFiles(opts.Src)
	if err != nil {
		return err
	}
	if len(src_files) == 0 {
		fmt.Println("No encrypted files")
	}
	for i := range src_files {
		src_file := src_files[i]
		ext := filepath.Ext(src_file)
		dest_file := src_file[:len(src_file)-len(ext)]
		fmt.Printf("Decrypt %s => %s\n", src_file, dest_file)
		err := DecryptFile(src_file, dest_file, opts.Password)
		if err != nil {
			fmt.Printf("Decrypt failed for %s", src_file)
			continue
		}
		os.Remove(src_file)
	}
	fmt.Printf("Decrypted %d files\n", len(src_files))
	return nil
}

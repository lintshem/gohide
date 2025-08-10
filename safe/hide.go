package safe

import (
	"fmt"
	"os"
)

func HideDir(opts Options) error {
	src_files, dest_files, err := GetFilteredPaths(opts.Src)
	if err != nil {
		fmt.Println("Can not read some files", err)
		return nil
	}
	for i := range src_files {
		src_file, dest_file := src_files[i], dest_files[i]
		err := EncryptFile(src_file, dest_file, opts.Password)
		if err != nil {
			fmt.Printf("Encryt failed for %s", src_file)
			continue
		}
		os.Remove(src_file)
	}
	fmt.Printf("Encrypted %d files\n", len(src_files))
	return nil
}

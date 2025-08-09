package safe

import (
	"fmt"
)

func ShowDir(opts Options) error {
	src_files, dest_files, err := GetFilteredPaths(opts.Src)
	if err != nil {
		fmt.Println("Can not read some files", err)
		return nil
	}
	for i := range src_files {
		src_file, dest_file := src_files[i], dest_files[i]
		EncryptFile(src_file, dest_file, opts.Password)
	}
	return nil
}

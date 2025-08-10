package safe

import (
	"fmt"
)

func CheckDir(opts Options) error {
	src_files, _, err := GetFilteredPaths(opts.Src)
	if err != nil {
		fmt.Println("Can not read some files", err)
		return nil
	}
	for _, file := range src_files {
		fmt.Printf("%s\n", file)
	}
	fmt.Printf("Checked %d files\n", len(src_files))
	return nil
}

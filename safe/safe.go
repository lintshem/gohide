package safe

import (
	"fmt"
	"log"
	"strings"
)

func Run(ops Options) error {
	handle_error := func(err error) {
		if err != nil {
			log.Fatal("Error: ", err)
		}
	}
	if strings.HasPrefix(ops.Mode, "e") {
		err := EncryptFile(ops.Src, ops.Dest, ops.Password)
		handle_error(err)
	} else if strings.HasPrefix(ops.Mode, "d") {
		err := DecryptFile(ops.Src, ops.Dest, ops.Password)
		handle_error(err)
	} else if strings.HasPrefix(ops.Mode, "h") {
		err := HideDir(ops)
		handle_error(err)
	} else if strings.HasPrefix(ops.Mode, "s") {
		err := ShowDir(ops)
		handle_error(err)
	} else if strings.HasPrefix(ops.Mode, "c") {
		err := CheckDir(ops)
		handle_error(err)
	} else if strings.HasPrefix(ops.Mode, "z") {
		err := ZipDir(ops)
		handle_error(err)
	} else {
		fmt.Println("Unknown mode:", ops.Mode)
	}

	return nil
}

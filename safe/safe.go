package safe

import (
	"fmt"
	"strings"
)

func Run(ops Options) error {

	if strings.HasPrefix(ops.Mode, "e") {
		err := EncryptFile(ops.Src, ops.Dest, ops.Password)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	} else if strings.HasPrefix(ops.Mode, "d") {
		err := DecryptFile(ops.Src, ops.Dest, ops.Password)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	} else if strings.HasPrefix(ops.Mode, "h") {
		err := HideDir(ops)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	} else if strings.HasPrefix(ops.Mode, "s") {
		err := ShowDir(ops)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	} else {
		fmt.Println("Unknown mode:", ops.Mode)
	}

	return nil
}

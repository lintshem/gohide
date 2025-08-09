package testfold

import (
	"fmt"
	"strings"
)

func Run(ops Options) error {
	if len(ops.Password) < 5 {
		fmt.Println("Error: Password too short!")
		return nil
	}
	if strings.HasPrefix(ops.Mode, "e") {
		if err := EncryptFile(ops.Src, ops.Dest, ops.Password); err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		fmt.Println("Encrypted to", ops.Dest)
	} else if strings.HasPrefix(ops.Mode, "d") {
		if err := DecryptFile(ops.Src, ops.Dest, ops.Password); err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		fmt.Println("Decrypted to", ops.Dest)
	} else {
		fmt.Println("Unknown mode:", ops.Mode)
	}
	return nil
}

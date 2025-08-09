package main

import (
	"fmt"
	"os"

	"com.secret.files/safe"
)

func main() {
	password := ""
	for {
		password = safe.ReadPassword()
		password2 := safe.ReadPassword()
		if len(password) < 5 {
			fmt.Println("Enter 5+ characters")
			continue
		}
		if password == password2 {
			break
		}
	}

	if len(os.Args) < 5 {
		fmt.Println("Usage: gohide <encrypt|decrypt> <source> <dest> <password>")
		return
	}

	ops := safe.Options{
		Mode:     os.Args[1],
		Src:      os.Args[2],
		Dest:     os.Args[3],
		Password: os.Args[4],
	}

	safe.Run(ops)

}

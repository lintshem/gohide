package main

import (
	"fmt"
	"os"

	"com.gosafe/safe"
)

func GetPassword() string {
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
	return password
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: gohide <hide|show|encrypt|decrypt> <source> <dest>")
		return
	}

	password := GetPassword()

	ops := safe.Options{
		Mode:     os.Args[1],
		Src:      os.Args[2],
		Dest:     os.Args[3],
		Password: password,
	}

	safe.Run(ops)

}

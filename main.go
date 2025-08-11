package main

import (
	"fmt"
	"os"
	"strings"

	"com.gosafe/safe"
)

func GetPassword() string {
	password := ""
	for {
		password = safe.ReadPassword("Enter Password: ")
		if len(password) < 3 {
			fmt.Println("Enter 3+ characters")
			continue
		}
		password2 := safe.ReadPassword("Confirm Password: ")
		if password == password2 {
			break
		} else {
			fmt.Println("Passwords din't match")
		}
	}
	return password
}

func GetOsArg(pos int, def string) string {
	if len(os.Args) <= pos {
		return def
	}
	return os.Args[pos]
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gohide <zip|hide|show|encrypt|decrypt> <source> <dest>")
		return
	}

	mode := strings.ToLower(os.Args[1])
	password := ""

	if !strings.HasPrefix(mode, "c") {
		password = GetPassword()
	}

	ops := safe.Options{
		Mode:     mode,
		Src:      GetOsArg(2, "."),
		Dest:     GetOsArg(3, "."),
		Password: password,
	}

	safe.Run(ops)

}

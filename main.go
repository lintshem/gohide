package main

import (
	"fmt"
	"os"

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
		fmt.Println("Usage: gohide <hide|show|encrypt|decrypt> <source> <dest>")
		return
	}

	password := GetPassword()

	ops := safe.Options{
		Mode:     os.Args[1],
		Src:      GetOsArg(2, "."),
		Dest:     GetOsArg(3, "."),
		Password: password,
	}

	safe.Run(ops)

}

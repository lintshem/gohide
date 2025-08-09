package safe

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/term"
)

func EncryptFile(srcPath, dstPath, password string) error {
	key := deriveKey(password)

	// Read file contents
	plain, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// Create nonce
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return err
	}

	// Encrypt
	encrypted := secretbox.Seal(nonce[:], plain, &nonce, key)

	// Write encrypted data
	return os.WriteFile(dstPath, encrypted, 0600)
}

func ReadPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
		panic("Could not read!")
	}
	fmt.Println()
	password := string(bytePassword)
	return password
}

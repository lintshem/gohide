package testfold

import (
	"crypto/rand"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
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

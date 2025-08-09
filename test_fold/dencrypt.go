package testfold

import (
	"errors"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func DecryptFile(srcPath, dstPath, password string) error {
	key := deriveKey(password)

	// Read encrypted data
	ciphertext, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}
	if len(ciphertext) < 24 {
		return errors.New("invalid file")
	}

	// Extract nonce
	var nonce [24]byte
	copy(nonce[:], ciphertext[:24])

	// Decrypt
	decrypted, ok := secretbox.Open(nil, ciphertext[24:], &nonce, key)
	if !ok {
		return errors.New("wrong password or corrupt file")
	}

	// Write decrypted data
	return os.WriteFile(dstPath, decrypted, 0600)
}

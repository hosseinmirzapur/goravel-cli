package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

/*
encrypt the content with passed key to the function
*/
func Encrypt(content, key []byte) ([]byte, error) {

	gcm, err := createGCMCipher(key)
	if err != nil {
		return nil, err
	}

	// encrypt content
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, content, nil)

	return cipherText, nil
}

/*
decrypt the content with passed key to the function
*/
func Decrypt(content, key []byte) ([]byte, error) {
	gcm, err := createGCMCipher(key)
	if err != nil {
		Error("decrypt", "error creating GCM cipher", err)
		os.Exit(1)
	}

	// decrypt content
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, content, nil)
}

// Creates GCM cipher
func createGCMCipher(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}

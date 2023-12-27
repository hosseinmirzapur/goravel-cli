package utils

import (
	"crypto/aes"
	"crypto/cipher"
	cryptRand "crypto/rand"
	"math/rand"
	"os"
)

type EncryptedData struct {
	EncryptedContent []byte
	Key              string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
encrypt the content with passed key to the function
*/
func Encrypt(content []byte) (*EncryptedData, error) {
	// generate key
	key := generateKey()

	gcm, err := createGCMCipher(key)
	if err != nil {
		return nil, err
	}

	// encrypt content
	nonce := make([]byte, gcm.NonceSize())
	if _, err := cryptRand.Read(nonce); err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, content, nil)

	return &EncryptedData{
		Key:              string(key),
		EncryptedContent: cipherText,
	}, nil
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
	nonceSize := gcm.NonceSize()
	if len(content) < nonceSize {
		Error("decrypt", "content is too short", nil)
		os.Exit(1)
	}
	nonce := content[:nonceSize]
	cipherText := content[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}

func generateKey() []byte {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
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

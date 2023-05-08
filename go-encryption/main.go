package main

import (
	"fmt"
	"log"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/md5"
	"io"
)

func createHash(key string) string {
	hasher:= md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, err := aes.NewCipher([]byte(createHash(passphrase)))
		if err != nil {
			log.Println(err.Error())
		}
	gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Println(err.Error())
		}

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
		if err != nil {
			log.Println(err.Error())
		}
	gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Println(err.Error())
		}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			log.Println(err.Error())
		}
	return plaintext
}

func main() {
	ciphertext := encrypt([]byte("Lets encrypt this boy!"), "password")
	fmt.Println(string(ciphertext))
	plaintext := decrypt(ciphertext, "password")
	fmt.Println(string(plaintext))
}
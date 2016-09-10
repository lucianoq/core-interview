package server

import (
	"core-interview/server/storage"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"errors"
	"encoding/base64"
)

func Store(id, plaintext string) (string, error) {
	aesKey, err := generateAesKey()
	if err != nil {
		return "", err
	}
	ciphertext, err := encrypt([]byte(plaintext), aesKey)
	if err != nil {
		return "", err
	}
	err = storage.Insert(id, toBase64(ciphertext))
	if err != nil {
		return "", err
	}
	return toBase64(aesKey), nil
}

func Retrieve(id, aesKey string) (string, error) {
	ciphertext, err := storage.Select(id)
	if err != nil {
		return "", err
	}
	cipherString, err := fromBase64(ciphertext)
	if err != nil {
		return "", err
	}
	keyString, err := fromBase64(aesKey)
	if err != nil {
		return "", err
	}
	plaintext, err := decrypt(cipherString, keyString)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func encrypt(text, key []byte) ([]byte, error) {
	var block cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Error creating Cipher")
	}
	ciphertext := make([]byte, aes.BlockSize + len(string(text)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.New("Error encrypting")
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return ciphertext, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	var block cipher.Block
	block, err := aes.NewCipher(key);
	if err != nil {
		return nil, errors.New("Error decrypting")
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("Error decrypting")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func toBase64(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

func fromBase64(in string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(in)
}

func generateAesKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
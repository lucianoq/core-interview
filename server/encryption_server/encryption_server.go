package server

import "core-interview/server/storage"

func Store(id, payload []byte) ([]byte, error) {
	aesKey := generateAesKey()
	cyphertext := encrypt(payload, aesKey)
	err := storage.Store(id, cyphertext)
	if err != nil {
		return nil, err
	}
	return aesKey, nil
}

func Retrieve(id, aesKey []byte) ([]byte, error) {
	cyphertext, err := storage.Retrieve(id)
	if err != nil {
		return nil, err
	}
	return decrypt(cyphertext, aesKey), nil
}

func encrypt(msg, key []byte) []byte {
	//TODO
	return msg
}

func decrypt(crypt, key []byte) []byte {
	//TODO
	return crypt
}

func generateAesKey() []byte {
	//TODO random
	return []byte("example")
}
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/tuanloc1105/go-common-lib/constant"
)

func GenerateCheckSumUsingSha256New(input string) (string, error) {
	if input == "" {
		return constant.EmptyString, errors.New("input can not be empty")
	}
	h := sha256.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GenerateCheckSumUsingSha256Sum256(input string) (string, error) {
	if input == "" {
		return constant.EmptyString, errors.New("input can not be empty")
	}
	sum := sha256.Sum256([]byte("this is a password"))
	return fmt.Sprintf("%x", sum), nil
}

func AesEncryption(stringToEncrypt string, keyString string) (encryptedString string, err error) {
	if stringToEncrypt == "" || keyString == "" {
		return constant.EmptyString, errors.New("input can not be empty")
	}
	if len(keyString) < 32 {
		return constant.EmptyString, errors.New("key length must be greater than 32 bytes")
	}

	//Since the key is in string, we need to convert decode it to bytes
	key, hexDecodeStringError := hex.DecodeString(keyString)
	if hexDecodeStringError != nil {
		return constant.EmptyString, hexDecodeStringError
	}
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, aesNewCipherError := aes.NewCipher(key)
	if aesNewCipherError != nil {
		return constant.EmptyString, aesNewCipherError
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, cipherNewGcmError := cipher.NewGCM(block)
	if cipherNewGcmError != nil {
		return constant.EmptyString, cipherNewGcmError
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, ioReadFullError := io.ReadFull(rand.Reader, nonce); ioReadFullError != nil {
		return constant.EmptyString, ioReadFullError
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func AesDecryption(encryptedString string, keyString string) (decryptedString string, err error) {
	if encryptedString == "" || keyString == "" {
		return constant.EmptyString, errors.New("input can not be empty")
	}
	if len(keyString) < 32 {
		return constant.EmptyString, errors.New("key length must be greater than 32 bytes")
	}

	key, keyDecodeStringError := hex.DecodeString(keyString)
	enc, encDecodeStringError := hex.DecodeString(encryptedString)
	if keyDecodeStringError != nil {
		return constant.EmptyString, keyDecodeStringError
	}
	if encDecodeStringError != nil {
		return constant.EmptyString, encDecodeStringError
	}

	//Create a new Cipher Block from the key
	block, aesNewCipherError := aes.NewCipher(key)
	if aesNewCipherError != nil {
		return constant.EmptyString, aesNewCipherError
	}

	//Create a new GCM
	aesGCM, cipherNewGcmError := cipher.NewGCM(block)
	if cipherNewGcmError != nil {
		return constant.EmptyString, cipherNewGcmError
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return constant.EmptyString, err
	}

	// return fmt.Sprintf("%s", plaintext), nil
	return string(plaintext), nil
}

func HashHMACSHA512(inputData, keyData []byte) []byte {
	formattedKey := formatKey(keyData)
	ipad := xorIPAD(formattedKey)
	innerData := Hash(append(ipad, inputData...))

	opad := xorOPAD(formattedKey)
	return Hash(append(opad, innerData...))
}

func formatKey(keyData []byte) []byte {
	if len(keyData) >= 128 {
		return Hash(keyData)
	}

	results := make([]byte, 128)
	copy(results, keyData)
	for i := len(keyData); i < 128; i++ {
		results[i] = 0x00
	}
	return results
}

func xorIPAD(k []byte) []byte {
	results := make([]byte, len(k))
	for i, key := range k {
		results[i] = key ^ 0x36
	}
	return results
}

func xorOPAD(k []byte) []byte {
	results := make([]byte, len(k))
	for i, key := range k {
		results[i] = key ^ 0x5C
	}
	return results
}

func Hash(inputValues []byte) []byte {
	formattedMessage := formatInput(inputValues)
	computed := compute(formattedMessage)
	return computed
}

func formatInput(inputValues []byte) []byte {
	// Implement your own formatting logic if needed
	return inputValues
}

func compute(messages []byte) []byte {
	// Implement your own computation logic if needed
	return messages
}

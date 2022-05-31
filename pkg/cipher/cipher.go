package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"io"
)

func EncryptAES(plainText, key []byte) (string, error) {
	byteKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	aesCipher, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data using aesGCM.Seal
	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data.
	// The first nonce argument in Seal is the prefix.
	return hex.EncodeToString(aesGCM.Seal(nonce, nonce, []byte(plainText), nil)), nil
}

func DecryptAES(cipherText, key string) (string, error) {
	byteKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	enc, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func EncryptRC4(plainText, key string) (string, error) {
	rc4Cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainText)
	dst := make([]byte, len(plainBytes))

	rc4Cipher.XORKeyStream(dst, plainBytes)

	return hex.EncodeToString(dst), nil
}

func DecryptRC4(cipherText, key string) (string, error) {
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	rc4Cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherBytes))
	rc4Cipher.XORKeyStream(plainText, cipherBytes)

	return string(plainText), nil
}

func GenerateKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"io"
)

func EncryptAES(plainText string, key []byte) (string, error) {
	aesCipher, err := aes.NewCipher(key)
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

// DecryptAES decrypts cipherText in HEX format and return plain text.
func DecryptAES(cipherText string, key []byte) (string, error) {
	enc, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", InvalidHexError{Message: err.Error()}
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	if len(enc) < nonceSize {
		return "", DecryptionError{Message: "nonce size cipher text is smaller than nonce size"}
	}

	// Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", DecryptionError{Message: err.Error()}
	}

	return string(plaintext), nil
}

func EncryptRC4(plainText string, key []byte) (string, error) {
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainText)
	dst := make([]byte, len(plainBytes))

	rc4Cipher.XORKeyStream(dst, plainBytes)

	return hex.EncodeToString(dst), nil
}

func DecryptRC4(cipherText string, key []byte) (string, error) {
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return "", err
	}

	enc, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", InvalidHexError{Message: err.Error()}
	}

	plainText := make([]byte, len(enc))
	rc4Cipher.XORKeyStream(plainText, enc)

	return string(plainText), nil
}

func GenerateKeyBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}

	return bytes, nil
}

func GenerateKeyHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

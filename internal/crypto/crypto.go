package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)

// scrypt params — tune for target machines. These are moderate.
// Increase N for stronger but slower derivation.
const (
	scryptN   = 1 << 15
	scryptR   = 8
	scryptP   = 1
	keyLen    = 32 // AES-256
	saltLen   = 16
	nonceLen  = 12 // AES-GCM nonce size
)

// GenerateRandomBytes returns securely random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	return b, err
}

// DeriveKey derives a key from masterPassword + salt using scrypt
func DeriveKey(masterPassword string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(masterPassword), salt, scryptN, scryptR, scryptP, keyLen)
}

// Encrypt returns base64(ciphertext), base64(salt), base64(nonce)
func Encrypt(plaintext []byte, masterPassword string) (string, string, string, error) {
	if masterPassword == "" {
		return "", "", "", errors.New("master password required")
	}
	salt, err := GenerateRandomBytes(saltLen)
	if err != nil {
		return "", "", "", err
	}

	key, err := DeriveKey(masterPassword, salt)
	if err != nil {
		return "", "", "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", "", err
	}

	nonce, err := GenerateRandomBytes(nonceLen)
	if err != nil {
		return "", "", "", err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	return base64.StdEncoding.EncodeToString(ciphertext),
		base64.StdEncoding.EncodeToString(salt),
		base64.StdEncoding.EncodeToString(nonce), nil
}

// Decrypt returns plaintext bytes given base64 inputs and masterPassword
func Decrypt(cipherB64, saltB64, nonceB64, masterPassword string) ([]byte, error) {
	if masterPassword == "" {
		return nil, errors.New("master password required")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return nil, err
	}
	salt, err := base64.StdEncoding.DecodeString(saltB64)
	if err != nil {
		return nil, err
	}
	nonce, err := base64.StdEncoding.DecodeString(nonceB64)
	if err != nil {
		return nil, err
	}

	key, err := DeriveKey(masterPassword, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		// authentication failed or wrong password
		return nil, errors.New("decryption failed: wrong master password or corrupted data")
	}
	return plaintext, nil
}

package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/config"
	"golang.org/x/crypto/argon2"
)

//TODO: Doesn't cater for numbers and special chars?

func GenerateSalt(length int32) ([]byte, error) {
	salt := make([]byte, length)

	_, err := rand.Read(salt)

	return salt, err
}

func GeneratePasswordHash(password string, salt []byte) []byte {
	return argon2.Key([]byte(password), salt, config.Auth.Time, config.Auth.Memory, config.Auth.Threads, config.Auth.KeyLen)
}

func GenerateEncodedSaltAndPasswordHash(password string) (encodedPassword string, encodedSalt string, err error) {
	salt, err := GenerateSalt(config.Auth.SaltLen)

	if err != nil {
		return encodedPassword, encodedSalt, err
	}

	hashedPassword := GeneratePasswordHash(password, salt)
	encodedSalt = hex.EncodeToString(salt)
	encodedPassword = hex.EncodeToString(hashedPassword)

	return encodedPassword, encodedSalt, err
}

func DecodeAndComparePasswords(plainTextPassword string, encodedPassword string, encodedSalt string) (bool, error) {
	salt, err := hex.DecodeString(encodedSalt)

	if err != nil {
		return false, err
	}

	fmt.Println("Salt", salt, encodedPassword)

	hashedPassword, err := hex.DecodeString(encodedPassword)

	if err != nil {
		return false, err
	}

	hashedComparison := GeneratePasswordHash(plainTextPassword, salt)

	return bytes.Equal(hashedComparison, hashedPassword), nil
}

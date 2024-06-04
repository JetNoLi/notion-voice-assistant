package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
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

	encodedSalt = base32.StdEncoding.EncodeToString(salt[:])
	encodedPassword = base32.StdEncoding.EncodeToString(hashedPassword[:])

	fmt.Printf("Password: %v\n", encodedPassword)
	fmt.Printf("Salt: %v\n", encodedSalt)

	return encodedPassword, encodedSalt, err
}

func DecodeAndComparePasswords(plainTextPassword string, encodedPassword string, encodedSalt string) (bool, error) {
	salt, err := base32.StdEncoding.DecodeString(encodedSalt)

	if err != nil {
		return false, err
	}

	hashedPassword, err := base32.StdEncoding.DecodeString(encodedPassword)

	if err != nil {
		return false, err
	}

	hashedComparison := GeneratePasswordHash(plainTextPassword, salt)

	encodedNewPassword := base32.StdEncoding.EncodeToString(hashedComparison)

	fmt.Println("Encoded New password", encodedNewPassword)
	fmt.Println("Encoded DB Value", encodedPassword)

	return bytes.Equal(hashedComparison, hashedPassword), nil
}

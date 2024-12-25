package password

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt генерирует случайную соль заданной длины
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// HashPassword хэширует пароль с использованием соли
func HashPassword(password string) (string, string, error) {
	// Генерация соли
	salt, err := GenerateSalt(16)
	if err != nil {
		return "", "", err
	}

	passwordWithSalt := password + salt

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return string(hash), salt, nil
}

// CheckPassword проверяет соответствие пароля, соли и хэша
func CheckPassword(password, salt, hash string) bool {
	passwordWithSalt := password + salt

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt))
	return err == nil
}

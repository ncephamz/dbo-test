package utils

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func StringToInt(value string) int {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)
}

func IntToString(value uint64) string {
	return strconv.FormatUint(uint64(value), 10)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

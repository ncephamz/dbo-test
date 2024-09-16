package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func StringToInt(value string) int {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)
}

func StringToUint64(value string) uint64 {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}

	return i
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

func GenerateID() uint64 {
	return uint64((time.Now().UnixNano() + int64(rand.Intn(100))))
}

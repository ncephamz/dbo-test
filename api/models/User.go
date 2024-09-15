package models

import (
	"errors"
	"html"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	stein "github.com/nasrul21/go-stein"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RandStringBytesMask(n int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 4
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return string(str)
}

func (u *User) Prepare() {
	u.ID = uuid.New().String()
	u.Phone = html.EscapeString(strings.TrimSpace(u.Phone))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Role = html.EscapeString(strings.TrimSpace(u.Role))
	u.Password = RandStringBytesMask(4)
	u.CreatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "token":
		if u.Phone == "" {
			return errors.New("Required Phone")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil

	default:
		if u.Phone == "" {
			return errors.New("Required Phone")
		}
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Role == "" {
			return errors.New("Required Role")
		}
		return nil
	}
}

func (u *User) SaveUser(c *stein.Stein) (*User, error) {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return &User{}, err
	}

	users := []User{
		{
			ID:        u.ID,
			Phone:     u.Phone,
			Name:      u.Name,
			Role:      u.Role,
			Password:  string(hashedPassword),
			CreatedAt: u.CreatedAt,
		},
	}
	_, _, err = c.Insert("users", users)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUsersByPhone(c *stein.Stein) ([]User, error) {
	var (
		result   []User
		filtered []User
	)

	readOption := stein.ReadOption{}

	_, err := c.Read("users", readOption, &result)
	if err != nil {
		return nil, err
	}

	for _, row := range result {
		if row.Phone == u.Phone {
			filtered = append(filtered, row)
		}
	}

	return filtered, nil
}

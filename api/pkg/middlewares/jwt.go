package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	Secret string
}

func (j Jwt) Signed(body interface{}, duration int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["body"] = body
	claims["exp"] = duration

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))

}

func (j Jwt) Validate(r *http.Request) error {
	tokenString := GetToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func GetToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func Pretty(data interface{}) interface{} {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil
	}

	return b
}

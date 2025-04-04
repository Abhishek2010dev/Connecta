package auth

import "github.com/golang-jwt/jwt/v5"

func GenerateJwt(claims Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateJwt(token string) (Claims, error) {}

package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(claims Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func DecodeJwtToken(tokenStr string, secret []byte) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		var tokenValidationErr error
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			tokenValidationErr = ErrTokenExpired
		case errors.Is(err, jwt.ErrInvalidType):
			tokenValidationErr = ErrInvalidTokenFormat
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			tokenValidationErr = ErrInvalidTokenSignature
		default:
			tokenValidationErr = ErrTokenValidationFailed
		}

		return nil, tokenValidationErr
	}

	if !token.Valid {
		return nil, ErrTokenValidationFailed
	}

	return &claims, nil
}

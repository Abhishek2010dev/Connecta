package auth

type TokenValidationError string

func (t TokenValidationError) Error() string {
	return string(t)
}

const (
	ErrTokenExpired          TokenValidationError = "Token has expired"
	ErrInvalidTokenFormat    TokenValidationError = "Invalid token format"
	ErrInvalidTokenSignature TokenValidationError = "Invalid token signature"
	ErrTokenValidationFailed TokenValidationError = "Token validation failed"
	RedisTokenNull           TokenValidationError = "Redis Token is null"
)

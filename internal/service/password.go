package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hash string) bool
}

type passwordServiceImpl struct{}

func NewPasswordService() Password {
	return &passwordServiceImpl{}
}

func (p *passwordServiceImpl) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}
	return string(hashPassword), nil
}

func (p *passwordServiceImpl) VerifyPassword(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

package service

import "github.com/Abhishek2010dev/Connecta/internal/dto"

type SessionService interface {
	GenerateToken(userId int64) (string, error)
	ValidateToken(token string) (*dto.AuthPaylaod, error)
}

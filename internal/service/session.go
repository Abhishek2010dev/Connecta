package service

import (
	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
)

type Session interface {
	GenerateToken(userId int64) (string, error)
	ValidateToken(token string) (*dto.AuthPaylaod, error)
}

type sessionServiceImpl struct {
	repo *repository.Session
}

func NewSession(repo *repository.Session) Session {
	return &sessionServiceImpl{repo}
}

func (s *sessionServiceImpl) GenerateToken(userId int64) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *sessionServiceImpl) ValidateToken(token string) (*dto.AuthPaylaod, error) {
	panic("not implemented") // TODO: Implement
}

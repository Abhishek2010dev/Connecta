package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
)

const (
	Day             = 24 * time.Hour
	SessionDuration = 30 * Day
	RenewBefore     = 15 * Day
)

type Session interface {
	GenerateToken(userID int64) (string, error)
	ValidateToken(token string) (*dto.AuthPaylaod, error)
}

type sessionServiceImpl struct {
	repo repository.Session
}

func NewSession(repo repository.Session) Session {
	return &sessionServiceImpl{repo}
}

func (s *sessionServiceImpl) GenerateToken(userID int64) (string, error) {
	bytes := make([]byte, 18)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	token := base64.RawURLEncoding.EncodeToString(bytes)
	if _, err := s.repo.Create(hashToken(token), userID, time.Now().Add(SessionDuration)); err != nil {
		return "", err
	}
	return token, nil
}

func (s *sessionServiceImpl) ValidateToken(token string) (*dto.AuthPaylaod, error) {
	panic("not implemented") // TODO: Implement
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
)

const (
	Day             = 24 * time.Hour
	SessionDuration = 30 * Day
	RenewBefore     = 15 * Day
)

type Session interface {
	GenerateToken(userID int64) (string, error)
	ValidateToken(token string) (*models.Session, *dto.AuthPaylaod, error)
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

func (s *sessionServiceImpl) ValidateToken(token string) (*models.Session, *dto.AuthPaylaod, error) {
	sessionID := hashToken(token)
	session, username, err := s.repo.FindByIDWithUsername(sessionID)
	if err != nil {
		return nil, nil, err
	}

	if time.Now().After(session.ExpiresAt) {
		log.Println("Session expired. Deleting...")
		if _, err := s.repo.DeleteByID(sessionID); err != nil {
			return nil, nil, err
		}
		return nil, nil, nil
	}

	if time.Now().After(session.ExpiresAt.Add(-RenewBefore)) {
		newExpiresAt := time.Now().Add(SessionDuration)
		session.ExpiresAt = newExpiresAt
		if _, err := s.repo.UpdateExpiration(sessionID, newExpiresAt); err != nil {
			return nil, nil, err
		}
	}

	return session, &dto.AuthPaylaod{
		UserId:   session.UserId,
		Username: username,
	}, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

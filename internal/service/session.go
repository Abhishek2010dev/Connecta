package service

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/internal/dto"
	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/internal/models"
	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/internal/repository"
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

func NewSession(db *sql.DB) Session {
	return &sessionServiceImpl{
		repo: repository.NewSessionRepository(db),
	}
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

var ErrSessionExpired = errors.New("session expired")

func (s *sessionServiceImpl) ValidateToken(token string) (*models.Session, *dto.AuthPaylaod, error) {
	sessionID := hashToken(token)
	session, username, err := s.repo.FindByIDWithUsername(sessionID)
	if err != nil {
		return nil, nil, err
	}

	if time.Now().After(session.ExpiresAt) {
		log.Printf("Session %s expired. Deleting...\n", sessionID)
		if _, err := s.repo.DeleteByID(sessionID); err != nil {
			return nil, nil, fmt.Errorf("failed to delete expired session: %w", err)
		}
		return nil, nil, ErrSessionExpired
	}

	if time.Now().After(session.ExpiresAt.Add(-RenewBefore)) {
		newExpiresAt := time.Now().Add(SessionDuration)
		session.ExpiresAt = newExpiresAt
		if _, err := s.repo.UpdateExpiration(sessionID, newExpiresAt); err != nil {
			return nil, nil, fmt.Errorf("failed to renew session expiration: %w", err)
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

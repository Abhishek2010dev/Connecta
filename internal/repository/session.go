package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/models"
)

// NOTE: I am returning sessionId in same function for just check
type SessionRepository interface {
	Create(tokenHash string, userID int64, expiresAt time.Time) (string, error)
	FindByIDWithUsername(sessionID string) (*models.Session, string, error)
	DeleteByID(sessionID string) (string, error)
	UpdateExpiration(sessionID string, newExpiration time.Time) (string, error)
}

type sessionRepoImpl struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepoImpl{db}
}

func (s *sessionRepoImpl) Create(tokenHash string, userID int64, expiresAt time.Time) (string, error) {
	query := "INSERT INTO session(id, user_id, expires_at) VALUES($1, $2, $3) RETURNING id"
	var sessionId string
	if err := s.db.QueryRow(query, tokenHash, userID, expiresAt).Scan(&sessionId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("Failed to create session: %w", err)
		}
		return "", fmt.Errorf("Failed to create session: %w", err)
	}
	return sessionId, nil
}

func (s *sessionRepoImpl) FindByIDWithUsername(sessionID string) (*models.Session, string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *sessionRepoImpl) DeleteByID(sessionID string) (string, error) {
	query := "DELETE FROM session WHERE id = $1 RETURNING id"
	var returnedID string
	if err := s.db.QueryRow(query, sessionID).Scan(&returnedID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("Failed to delete session with id %s: %w", sessionID, err)
	}
	return returnedID, nil
}

func (s *sessionRepoImpl) UpdateExpiration(sessionID string, newExpiration time.Time) (string, error) {
	panic("not implemented") // TODO: Implement
}

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
	FindByIDWithUsername(sessionID string) (*models.Session, int64, error)
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

func (s *sessionRepoImpl) FindByIDWithUsername(sessionID string) (*models.Session, int64, error) {
	query := `
		SELECT s.id, s.user_id, s.expires_at, u.username
	 	FROM session s
		JOIN users u ON s.user_id = u.id
		WHERE s.id = $1
	`

	var session models.Session
	var userId int64
	err := s.db.QueryRow(query, sessionID).Scan(&session.Id, &session.UserId, &session.ExpiresAt, &userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, nil
		}
		return nil, 0, fmt.Errorf("Failed to update session with id %s: %w", sessionID, err)
	}

	return &session, userId, nil
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
	query := "UPDATE  session SET expires_at = $1 WHERE id = $2 RETURNING id"
	var returnedID string
	if err := s.db.QueryRow(query, newExpiration, sessionID).Scan(&returnedID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("Failed to update session with id %s: %w", sessionID, err)
	}
	return returnedID, nil
}

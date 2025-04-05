package repository

import (
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type SessionRepository interface {
	// Create creates a new session for a given user ID and token.
	Create(token string, userID int64) (*models.Session, error)

	// FindByIDWithUsername fetches a session by its ID and returns the session and associated username.
	FindByIDWithUsername(sessionID string) (*models.Session, string, error)

	// DeleteByID deletes a session by its ID and returns the deleted session ID or a confirmation message.
	DeleteByID(sessionID string) (string, error)

	// UpdateExpiration updates the expiration time of a session.
	UpdateExpiration(sessionID string, newExpiration time.Time) error
}

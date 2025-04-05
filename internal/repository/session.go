package repository

import (
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/models"
)

// NOTE: I am returning sessionId in same function for just check
type SessionRepository interface {
	Create(token string, userID int64) (string, error)
	FindByIDWithUsername(sessionID string) (*models.Session, string, error)
	DeleteByID(sessionID string) (string, error)
	UpdateExpiration(sessionID string, newExpiration time.Time) (string, error)
}

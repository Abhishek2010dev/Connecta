package repository

import (
	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type UserRespoitory interface {
	Create(payload dto.CreateUserPayload) (*models.User, error)
	ExitsByEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}

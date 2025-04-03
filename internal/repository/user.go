package repository

import (
	"database/sql"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type UserRepository interface {
	Create(payload dto.CreateUserPayload) (*models.User, error)
	ExitsByEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUser(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) Create(payload dto.CreateUserPayload) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userRepositoryImpl) ExitsByEmail(email string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

package repository

import (
	"database/sql"
	"fmt"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type UserRepository interface {
	Create(payload dto.CreateUserPayload) (int64, error)
	ExitsByEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUser(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) Create(payload dto.CreateUserPayload) (int64, error) {
	query := `
		INSERT INTO users(name, username, email, password) VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userId int64
	err := u.db.QueryRow(query, payload.Name, payload.Username, payload.Email, payload.Password).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("Failed to create user: %w", err)
		}
		return 0, fmt.Errorf("Failed to scan userId: %w", err)
	}

	return userId, nil
}

func (u *userRepositoryImpl) ExitsByEmail(email string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

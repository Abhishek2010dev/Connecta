package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type User interface {
	Create(payload dto.CreateUserPayload) (int64, error)
	ExistsByEmailAndUsername(email string, username string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return &userRepoImpl{db}
}

func (u *userRepoImpl) Create(payload dto.CreateUserPayload) (int64, error) {
	query := `
		INSERT INTO users(name, username, email, password) VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userId int64
	err := u.db.QueryRow(query, payload.Name, payload.Username, payload.Email, payload.Password).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("Failed to create user: %w", err)
		}
		return 0, fmt.Errorf("Failed to scan userId: %w", err)
	}

	return userId, nil
}

func (u *userRepoImpl) ExistsByEmailAndUsername(email string, username string) (bool, error) {
	query := "SELECT 1 FROM users WHERE email = $1 AND username = $2"
	var exits int
	if err := u.db.QueryRow(query, email, username).Scan(&exits); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("Failed to check email existence: %w", err)
	}
	return true, nil
}

func (u *userRepoImpl) FindByEmail(email string) (*models.User, error) {
	query := `
        	SELECT id, name, username, email, password, created_at FROM users 
        	WHERE email = $1
	`

	var user models.User
	row := u.db.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("User does not exits: %w", err)
		}
		return nil, fmt.Errorf("Failed to find user by email %s: %w", email, err)
	}
	return &user, nil
}

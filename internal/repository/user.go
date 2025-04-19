package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/models"
)

type User interface {
	Create(payload *dto.CreateUserPayload) (int64, error)
	CheckEmailAndUsername(email, username string) (emailExists bool, usernameExists bool, err error)
	FindByEmail(email string) (*models.User, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return &userRepoImpl{db}
}

func (u *userRepoImpl) Create(payload *dto.CreateUserPayload) (int64, error) {
	query := `
		INSERT INTO users(name, username, email, password) VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userId int64
	err := u.db.QueryRow(query, payload.Name, payload.Username, strings.ToLower(payload.Email), payload.Password).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("Failed to create user: %w", err)
		}
	}

	return userId, nil
}

func (u *userRepoImpl) CheckEmailAndUsername(email, username string) (emailExists bool, usernameExists bool, err error) {
	queryEmail := "SELECT 1 FROM users WHERE email = $1"
	if err := u.db.QueryRow(queryEmail, email).Scan(&emailExists); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, false, fmt.Errorf("Failed to check email existence: %w", err)
	}

	queryUsername := "SELECT 1 FROM users WHERE username = $1"
	if err := u.db.QueryRow(queryUsername, username).Scan(&usernameExists); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, false, fmt.Errorf("Failed to check username existence: %w", err)
	}

	return
}

func (u *userRepoImpl) ExistsByUsername(username string) (bool, error) {
	query := "SELECT 1 FROM users WHERE username = $1"
	var exists int
	if err := u.db.QueryRow(query, username).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("Failed to check username existence: %w", err)
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
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to find user by email %s: %w", email, err)
	}
	return &user, nil
}

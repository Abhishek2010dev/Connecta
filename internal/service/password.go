package service

type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hash string) bool
}

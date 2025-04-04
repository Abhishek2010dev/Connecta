package service

type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hash string) bool
}

type passwordServiceImpl struct{}

func NewPasswordService() PasswordService {
	return &passwordServiceImpl{}
}

func (p *passwordServiceImpl) HashPassword(password string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (p *passwordServiceImpl) VerifyPassword(password string, hash string) bool {
	panic("not implemented") // TODO: Implement
}

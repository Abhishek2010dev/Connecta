package config

type Auth struct {
	JwtSecret string
}

func NewAuth() Auth {
	return Auth{
		JwtSecret: LoadEnv("AUTH_JWT_SECRET"),
	}
}

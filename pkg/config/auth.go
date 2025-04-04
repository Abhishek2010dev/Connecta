package config

type Auth struct {
	JwtSecret string
}

func NewAuth() Auth {
	return Auth{
		JwtSecret: LoadEnv("Auth_Jwt_Secret"),
	}
}

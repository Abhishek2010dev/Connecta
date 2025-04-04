package config

type Auth struct {
	SessionSecret string
}

func NewAuth() Auth {
	return Auth{
		SessionSecret: LoadEnv("AUTH_SESSION_SECRET"),
	}
}

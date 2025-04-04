package config

type Auth struct {
	AccessSecret  string
	RefreshSecret string
}

func NewAuth() Auth {
	return Auth{
		AccessSecret:  LoadEnv("AUTH_ACCESS_SECRET"),
		RefreshSecret: LoadEnv("AUTH_REFRESH_TOKEN"),
	}
}

package config

type Database struct {
	URL string
}

func NewDatabase() Database {
	return Database{
		URL: getEnv("DATABASE_URL"),
	}
}

package models

import "time"

type User struct {
	Id        int
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

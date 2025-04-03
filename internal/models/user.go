package models

import "time"

type User struct {
	id         int
	name       string
	username   string
	email      string
	password   string
	created_at time.Time
}

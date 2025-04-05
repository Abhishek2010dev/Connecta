package models

import "time"

type Session struct {
	Id        string
	UserId    int64
	ExpiresAt time.Time
}

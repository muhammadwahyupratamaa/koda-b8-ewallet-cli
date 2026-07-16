package model

import "time"

type UserSession struct {
	ID           int64
	UserID       int64
	SessionToken string
	ExpiredAt    time.Time
}
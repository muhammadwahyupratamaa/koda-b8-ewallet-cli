package model

import "time"

type PasswordReset struct {
	ID         int
	UserID     int
	Token string
	ExpiredAt  time.Time
	CreatedAt  time.Time
}
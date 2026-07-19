package model

import "time"

type PasswordReset struct {
	ID         int
	UserID     int
	ResetToken string
	ExpiredAt  time.Time
	CreatedAt  time.Time
}
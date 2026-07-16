package model

import "time"

type User struct {
	ID        int64
	FullName  string
	UserName  string
	Age       int
	Address   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
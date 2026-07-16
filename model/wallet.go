package model

import "time"

type Wallet struct {
	ID int64
	userID int64
	walletNumber string
	balance float64
	status string
	createdAt time.Time
	updatedAt time.Time
}
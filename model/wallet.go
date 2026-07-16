package model

import "time"

type Wallet struct {
	ID int64
	UserID int64
	WalletNumber string
	Balance float64
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}
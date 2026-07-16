package model

import "time"

type Transaction struct{
	ID int64
	SenderWalletID int64
	ReceiverWalletID int64
	Amount float64
	Type string
	Status string
	Description string
	CreatedAt time.Time
}
package repository

import (
	"context"
	"koda-b8-ewallet-cli/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository struct{
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) *WalletRepository{
return  &WalletRepository{
	db: db,
}
}
func (r *WalletRepository) GetWalletByUserID(userID int64 )(*model.Wallet, error) {
	var wallet model.Wallet
	
	err := r.db.QueryRow(context.Background(),`
	SELECT ID,user_id,wallet_number,balance,status,created_at,updated_At FROM wallets WHERE user_id=$1`,userID).Scan(&wallet.ID,&wallet.UserID,&wallet.WalletNumber,&wallet.Balance,&wallet.Status,&wallet.CreatedAt,&wallet.UpdatedAt,)
	if err != nil {
		return  nil, err
	}
	return &wallet , nil
}
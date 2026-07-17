package repository

import (
	"context"
	"errors"
	"koda-b8-ewallet-cli/internal/model"

	"github.com/jackc/pgx/v5"
)

type WalletRepository struct{
	db *pgx.Conn
}

func NewWalletRepository(db *pgx.Conn) *WalletRepository{
return  &WalletRepository{
	db: db,
}
}
func (r *WalletRepository) GetWalletByUserID(userID int64 )(*model.Wallet, error) {
	var wallet model.Wallet
	
	err := r.db.QueryRow(context.Background(),`
	SELECT ID,user_id,wallet_number,balance,status,created_at,updated_At FROM wallets WHERE user_id=$1`,userID).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.WalletNumber,
		&wallet.Balance,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,)
	if err != nil {
		return  nil, err
	}
	return &wallet , nil
}

func (r *WalletRepository) GetWalletByWalletNumber(walletNumber string) (*model.Wallet, error){
	var wallet model.Wallet

	err:= r.db.QueryRow(context.Background(),`
	SELECT ID,user_id,wallet_number,balance,status,created_at,updated_at FROM wallets WHERE wallet_number=$1`,walletNumber).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.WalletNumber,
		&wallet.Balance,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		return nil , err
	}
	return &wallet , nil
}

func (r *WalletRepository) UpdateBalance(walletID int64, balance int) error {
	result, err := r.db.Exec(
		context.Background(),
		`
		UPDATE wallets
		SET
			balance = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		balance,
		walletID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("wallet not found")
	}

	return nil
}

func (r *WalletRepository) CreateWallet(userID int64, walletNumber string) error {
	_, err := r.db.Exec(
		context.Background(),
		`
		INSERT INTO wallets (
			user_id,
			wallet_number,
			balance,
			status
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		`,
		userID,
		walletNumber,
		0,
		"ACTIVE",
	)

	return err
}
package repository

import (
	"context"
	"koda-b8-ewallet-cli/internal/model"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	_, err := r.db.Exec(
		context.Background(),
		`
		INSERT INTO transactions (
			sender_wallet_id,
			receiver_wallet_id,
			amount,
			type,
			status,
			description
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		`,
		transaction.SenderWalletID,
		transaction.ReceiverWalletID,
		transaction.Amount,
		transaction.Type,
		transaction.Status,
		transaction.Description,
	)

	if err != nil {
		return err
	}

	return nil
}
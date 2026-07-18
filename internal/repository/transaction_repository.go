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

func (r *TransactionRepository) GetTransactionHistory(walletID int) ([]model.TransactionHistory, error) {
	query := `
	SELECT
		t.id,
		t.sender_wallet_id,
		t.receiver_wallet_id,
		sw.wallet_number,
		rw.wallet_number,
		t.amount,
		t.status,
		t.created_at
	FROM transactions t
	JOIN wallets sw ON sw.id = t.sender_wallet_id
	JOIN wallets rw ON rw.id = t.receiver_wallet_id
	WHERE t.sender_wallet_id = $1
	   OR t.receiver_wallet_id = $1
	ORDER BY t.created_at DESC;
	`

	rows, err := r.db.Query(context.Background(), query, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []model.TransactionHistory

	for rows.Next() {
		var history model.TransactionHistory

		err := rows.Scan(
			&history.ID,
			&history.SenderWalletID,
			&history.ReceiverWalletID,
			&history.SenderWalletNumber,
			&history.ReceiverWalletNumber,
			&history.Amount,
			&history.Status,
			&history.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		histories = append(histories, history)
	}

	return histories, nil
}
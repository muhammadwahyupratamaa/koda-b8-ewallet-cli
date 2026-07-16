package service

import (
	"context"
	"errors"

	"koda-b8-ewallet-cli/internal/repository"
	"github.com/jackc/pgx/v5"
)

type TransferService struct {
	db              *pgx.Conn
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepository
}

func NewTransferService(
	db *pgx.Conn,
	walletRepo *repository.WalletRepository,
	transactionRepo *repository.TransactionRepository,
) *TransferService {
	return &TransferService{
		db:              db,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *TransferService) Transfer(
	senderWalletNumber string,
	receiverWalletNumber string,
	amount int,
) error {

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	sender, err := s.walletRepo.GetWalletByWalletNumber(senderWalletNumber)
	if err != nil {
		return err
	}

	receiver, err := s.walletRepo.GetWalletByWalletNumber(receiverWalletNumber)
	if err != nil {
		return err
	}

	if sender.Balance < amount {
		return errors.New("insufficient balance")
	}

	_, err = tx.Exec(
		context.Background(),
		`
		UPDATE wallets
		SET
			balance = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		sender.Balance-amount,
		sender.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		`
		UPDATE wallets
		SET
			balance = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		receiver.Balance+amount,
		receiver.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
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
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
		`,
		sender.ID,
		receiver.ID,
		amount,
		"TRANSFER",
		"SUCCESS",
		"Transfer Success",
	)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
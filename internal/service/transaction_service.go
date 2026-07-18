package service

import (
	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/repository"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	walletRepo      *repository.WalletRepository
}

func NewTransactionService(
	transactionRepo *repository.TransactionRepository,
	walletRepo *repository.WalletRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
	}
}

func (s *TransactionService) GetHistory(userID int) (int, []model.TransactionHistory, error) {

	wallet, err := s.walletRepo.GetWalletByUserID(int64(userID))
	if err != nil {
		return 0, nil, err
	}

	histories, err := s.transactionRepo.GetTransactionHistory(int(wallet.ID))
	if err != nil {
		return 0, nil, err
	}

	return int(wallet.ID), histories, nil
}
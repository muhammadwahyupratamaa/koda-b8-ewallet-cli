package main

import (
	"koda-b8-ewallet-cli/internal/database"
	"koda-b8-ewallet-cli/internal/menu"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/service"
)

func main() {

	db := database.Connect()
	defer db.Close()

	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	_ = service.NewTransferService(
		db,
		walletRepo,
		transactionRepo,
	)

	menu.MainMenu()

}
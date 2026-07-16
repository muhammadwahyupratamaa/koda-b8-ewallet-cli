package main

import (
	"context"
	"koda-b8-ewallet-cli/internal/database"
	"koda-b8-ewallet-cli/internal/menu"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/service"
)

func main() {
	
	db := database.Connect()
	defer db.Close(context.Background())
	sessionRepo := repository.NewSessionRepository(db)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(
	userRepo,
	sessionRepo,
	)
	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	walletService := service.NewWalletService(walletRepo)


	_ = service.NewTransferService(
		db,
		walletRepo,
		transactionRepo,
	)

	menu.MainMenu(authService, walletService)

}
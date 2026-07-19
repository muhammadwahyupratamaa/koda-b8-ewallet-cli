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
	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)
	authService := service.NewAuthService(
	userRepo,
	sessionRepo,
	walletRepo,
	)

	passwordResetService := service.NewPasswordResetService(
	userRepo,
	passwordResetRepo,
)
	walletService := service.NewWalletService(walletRepo)


	transferService := service.NewTransferService(
		db,
		walletRepo,
		transactionRepo,
	)
	transactionService := service.NewTransactionService(
	transactionRepo,
	walletRepo,
)

	menu.MainMenu(authService, walletService,transferService,transactionService,passwordResetService)

}
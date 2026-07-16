package main

import (
	"fmt"
	"log"

	"koda-b8-ewallet-cli/internal/database"
	"koda-b8-ewallet-cli/internal/repository"
	"koda-b8-ewallet-cli/internal/service"
)

func main() {
	db := database.Connect()
	defer db.Close()

	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	transferService := service.NewTransferService(
		db,
		walletRepo,
		transactionRepo,
	)

	// err := transferService.Transfer(
	// 	"WAL000001",
	// 	"WAL000002",
	// 	100000,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Transfer Success")
}
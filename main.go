package main

import (
	"fmt"
	"log"

	"koda-b8-ewallet-cli/internal/database"
	"koda-b8-ewallet-cli/internal/repository"
)

func main() {
	db := database.Connect()
	defer db.Close()

	walletRepo := repository.NewWalletRepository(db)

	wallet, err := walletRepo.GetWalletByUserID(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===== WALLET =====")
	fmt.Println("ID            :", wallet.ID)
	fmt.Println("User ID       :", wallet.UserID)
	fmt.Println("Wallet Number :", wallet.WalletNumber)
	fmt.Println("Balance       :", wallet.Balance)
	fmt.Println("Status        :", wallet.Status)
	fmt.Println("Created At    :", wallet.CreatedAt)
	fmt.Println("Updated At    :", wallet.UpdatedAt)
}
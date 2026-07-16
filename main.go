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

err := walletRepo.UpdateBalance(999, 750000)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Update balance berhasil!")
}
package menu

import (
	"fmt"

	"koda-b8-ewallet-cli/internal/service"
)

func ShowLoginMenu(
	authService *service.AuthService,
	walletService *service.WalletService,
	transferService *service.TransferService,
	transactionService *service.TransactionService,
) {

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("          LOGIN")
	fmt.Println("==============================")

	fmt.Print("Username : ")
	username := Input()

	fmt.Print("Password : ")
	password := Input()

	user, err := authService.Login(
	username,
	password,
	)

	if err != nil {
	fmt.Println("Login Failed:", err)
	Pause()
	return
	}

	fmt.Println("Login Success")
	Pause()

	Dashboard(user, walletService,transferService,transactionService)
}
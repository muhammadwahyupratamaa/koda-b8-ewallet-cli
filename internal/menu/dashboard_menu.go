package menu

import (
	"fmt"
	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/service"
	"os"
)

func ShowDashboard(user *model.User) {

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("          GOPAY MENU")
	fmt.Println("===================================")
	fmt.Printf("Welcome, %s\n", user.UserName)
	fmt.Println("===================================")
	fmt.Println("1. Show Balance")
	fmt.Println("2. Transfer")
	fmt.Println("3. History")
	fmt.Println("4. Logout")
	fmt.Println("0. Exit")
	fmt.Println("===================================")
	fmt.Print("Choose Menu : ")
}

func Dashboard(
    user *model.User,
    walletService *service.WalletService,
    transferService *service.TransferService,
) {
    for {
        ShowDashboard(user)
        choice := Input()

        switch choice {
        case "1": wallet, err := walletService.ShowBalance(user.ID)
			 if err != nil {
			 fmt.Println(err)
			 break
			}

			fmt.Println()
			fmt.Println("===== WALLET =====")
			fmt.Println("Wallet Number :", wallet.WalletNumber)
			fmt.Println("Balance       :", wallet.Balance)
        case "2":
			ShowTransferMenu(user,walletService,transferService)
        case "3":
        case "4": fmt.Println("Logout berhasil")
            return

        case "0":
            os.Exit(0)
        }
    }
}
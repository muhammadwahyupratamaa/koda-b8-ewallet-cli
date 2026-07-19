package menu

import (
	"fmt"
	"strconv"

	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/service"
)

func ShowTransferMenu(
    user *model.User,
    walletService *service.WalletService,
    transferService *service.TransferService,
) {

	fmt.Println()
	fmt.Println("===== TRANSFER =====")

	fmt.Print("Destination Wallet : ")
	receiverWallet := Input()

	fmt.Print("Amount : ")
	amountStr := Input()

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
	fmt.Println("Invalid amount")
	return
	}

	wallet, err := walletService.ShowBalance(user.ID)
	if err != nil {
	fmt.Println(err)
	Pause()
	return
	}
	fmt.Println(wallet.WalletNumber)


	err = transferService.Transfer(
	wallet.WalletNumber,
	receiverWallet,
	amount,
	)

	if err != nil {
	fmt.Println("Transfer Failed:", err)
	return
	}
	fmt.Println("Transfer Success!")
	Pause()
}
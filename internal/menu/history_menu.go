package menu

import (
	"fmt"

	"koda-b8-ewallet-cli/internal/model"
	"koda-b8-ewallet-cli/internal/service"
)

func ShowTransferHistory(user *model.User, transactionService *service.TransactionService) {

	walletID, histories, err := transactionService.GetHistory(int(user.ID))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n===== TRANSFER HISTORY =====")

	if len(histories) == 0 {
		fmt.Println("No transaction history.")
		Pause()
		return
	}

	for i, h := range histories {

		transactionType := "RECEIVED"
		target := h.SenderWalletNumber

		if h.SenderWalletID == walletID {
			transactionType = "SENT"
			target = h.ReceiverWalletNumber
		}

		fmt.Printf(
			"%d. %-8s %-12s Rp %d (%s)\n",
			i+1,
			transactionType,
			target,
			h.Amount,
			h.Status,
		)
	}
	Pause()
}
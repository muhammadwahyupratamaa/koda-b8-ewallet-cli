package menu

import (
	"fmt"
	"koda-b8-ewallet-cli/internal/service"
)

func MainMenu(authService *service.AuthService) {

	for {

		ShowMainMenu()

		choice := Input()

		switch choice {
		case "1":
			ShowRegisterMenu(authService)
		case "2":
			ShowLoginMenu(authService)
		case "0":
			fmt.Println("Thank you...")
			return
		default:
			fmt.Println("Invalid Menu")
		}
	}
}
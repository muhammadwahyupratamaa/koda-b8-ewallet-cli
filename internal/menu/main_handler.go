package menu

import "fmt"

func MainMenu() {

	for {

		ShowMainMenu()

		choice := Input()

		switch choice {
		case "1":
			ShowRegisterMenu()
		case "2":
			fmt.Println("Login")
		case "0":
			fmt.Println("Thank you...")
			return
		default:
			fmt.Println("Invalid Menu")
		}
	}
}
package menu

import (
	"fmt"

	"koda-b8-ewallet-cli/internal/service"
)

func ShowLoginMenu(authService *service.AuthService) {

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
	return
	}

	fmt.Println("Login Success")

	Dashboard(user.UserName)
}
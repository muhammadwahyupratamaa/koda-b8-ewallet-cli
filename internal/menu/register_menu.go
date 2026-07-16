package menu

import (
	"fmt"
	"koda-b8-ewallet-cli/internal/service"
	"strconv"
)

func ShowRegisterMenu(authService *service.AuthService) {

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("         REGISTER")
	fmt.Println("==============================")

	fmt.Print("Full Name : ")
	fullName := Input()

	fmt.Print("Username  : ")
	username := Input()

	fmt.Print("Age       : ")
	ageInput := Input()

	age, err := strconv.Atoi(ageInput)
	if err != nil {
	fmt.Println("Age must be a number")
	return
	}

	fmt.Print("Address   : ")
	address := Input()

	fmt.Print("Email     : ")
	email := Input()

	fmt.Print("Password  : ")
	password := Input()

	fmt.Println()
	fmt.Println("===== REGISTER DATA =====")
	fmt.Println("Full Name :", fullName)
	fmt.Println("Username  :", username)
	fmt.Println("Age       :", age)
	fmt.Println("Address   :", address)
	fmt.Println("Email     :", email)
	fmt.Println("Password  :", password)
	fmt.Println("=========================")

	err = authService.Register(
	fullName,
	username,
	age,
	address,
	email,
	password,
	)	

	if err != nil {
	fmt.Println("Register Failed:", err)
	return
	}

	fmt.Println("Register Success")
}
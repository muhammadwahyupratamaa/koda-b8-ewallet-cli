package menu

import "fmt"

func ShowRegisterMenu() {

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("         REGISTER")
	fmt.Println("==============================")

	fmt.Print("Full Name : ")
	fullName := Input()

	fmt.Print("Username  : ")
	username := Input()

	fmt.Print("Age       : ")
	age := Input()

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
}
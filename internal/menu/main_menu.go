package menu

import "fmt"

func ShowMainMenu() {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("          E-WALLET CLI")
	fmt.Println("===================================")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("0. Exit")
	fmt.Println("===================================")
	fmt.Print("Choose Menu : ")
}
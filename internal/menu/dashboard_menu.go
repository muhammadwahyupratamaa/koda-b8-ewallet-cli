package menu

import (
	"fmt"
	"os"
)

func ShowDashboard(username string) {

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("          GOPAY MENU")
	fmt.Println("===================================")
	fmt.Printf("Welcome, %s\n", username)
	fmt.Println("===================================")
	fmt.Println("1. Show Balance")
	fmt.Println("2. Transfer")
	fmt.Println("3. Transfer History")
	fmt.Println("4. Logout")
	fmt.Println("0. Exit")
	fmt.Println("===================================")
	fmt.Print("Choose Menu : ")
}

func Dashboard(username string) {

    for {
        ShowDashboard(username)
        choice := Input()

        switch choice {
        case "1":
        case "2":
        case "3":
        case "4":
            return

        case "0":
            os.Exit(0)
        }
    }
}
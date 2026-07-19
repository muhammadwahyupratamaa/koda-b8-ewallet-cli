package menu

import (
	"fmt"

	"koda-b8-ewallet-cli/internal/service"
)

func ShowForgotPasswordMenu(
	passwordResetService *service.PasswordResetService,
) {

	fmt.Println()
	fmt.Println("===== FORGOT PASSWORD =====")

	fmt.Print("Username : ")
	username := Input()

	token, err := passwordResetService.GenerateResetToken(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Reset Token :", token)

	fmt.Print("Enter Token : ")
	inputToken := Input()

	fmt.Print("New Password : ")
	newPassword := Input()

	fmt.Print("Confirm Password : ")
	confirmPassword := Input()

	if newPassword != confirmPassword {
		fmt.Println("Password confirmation does not match")
		return
	}

	err = passwordResetService.ResetPassword(
		inputToken,
		newPassword,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Password updated successfully.")
}
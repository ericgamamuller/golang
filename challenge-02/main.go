package main

import (
	"fmt"

	"git/challenge-02/config"
	"git/challenge-02/services"
)

const version = "0.1"

var payments []payment

// Type

type payment struct {
	PaymentId          string
	Status             int
	ConfirmationStatus int
	Voids              []void
}

type void struct {
	VoidId     string
	PaymentId  string
	VoidStatus int
}

func main() {
	config.Initialize()

	welcome()

	menu()

	goodbye()
}

// Ornament

func welcome() {
	fmt.Println("----- GO Lang : Challenge-02 -----")
	fmt.Println("- Version:", version)
	fmt.Println("Welcome!")
	fmt.Println()
}

func goodbye() {
	fmt.Println("Thanks for using this program! See you next time!")
}

// Menus

func menu() {
loop:
	for {
		fmt.Println("Please inform the operation you would like to execute:")
		fmt.Println("  1 - Execute new payment")
		fmt.Println("  2 - Select existing payment")
		fmt.Println("  0 - End Program")
		fmt.Println()

		var option int
		fmt.Print("Option: ")
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case 1:
			newPayment()
			menuPayment(len(payments) - 1)
		case 2:
			if len(payments) > 0 {
				menuPayment(selectPayment())
			} else {
				fmt.Println("No previous payment to select")
			}
		case 0:
			break loop
		default:
			fmt.Println(option, "is not a valid option")
		}
		fmt.Println()
	}
}

func selectPayment() int {
	fmt.Println("Select a Payment:")
	for i, p := range payments {
		fmt.Println(" ", i, "-", p.PaymentId)
	}
	fmt.Println()

	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)
	fmt.Println()

	if option >= len(payments) {
		fmt.Println("This payment doesn't exists. Selecting latest as default")
		option = len(payments) - 1
	}

	return option
}

func menuPayment(index int) {
loopPayment:
	for {
		fmt.Println("Selected Payment:")
		fmt.Println("    PaymentId:", payments[index].PaymentId)
		fmt.Println("    Status:", payments[index].Status)
		fmt.Println("    ConfirmationStatus:", payments[index].ConfirmationStatus)
		fmt.Println()
		fmt.Println("Please inform the operation you would like to execute:")
		fmt.Println("  1 - Reverse Payment")
		fmt.Println("  2 - Confirm Payment")
		fmt.Println("  0 - Go back")
		fmt.Println()

		var option int
		fmt.Print("Option: ")
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case 1:
			reversePayment(index)
		case 2:
			confirmPayment(index)
		case 0:
			break loopPayment
		default:
			fmt.Println(option, "is not a valid option")
		}
		fmt.Println()
	}
}

// Execute

func newPayment() {
	var amount int
	fmt.Print("Inform new payment amount: ")
	fmt.Scan(&amount)

	response := services.Payment(amount)
	fmt.Println("Payment Response:")
	fmt.Println("    PaymentId:", response.Payment.PaymentId)
	fmt.Println("    Status:", response.Payment.Status)
	fmt.Println("    ConfirmationStatus:", response.Payment.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.Payment.ReturnCode)
	fmt.Println("    ReturnMessage:", response.Payment.ReturnMessage)
	fmt.Println()

	newPayment := payment{
		PaymentId:          response.Payment.PaymentId,
		Status:             response.Payment.Status,
		ConfirmationStatus: response.Payment.ConfirmationStatus,
	}
	payments = append(payments, newPayment)
}

func reversePayment(index int) {
	response := services.PaymentReverse(payments[index].PaymentId)
	fmt.Println("Payment Reversal Response:")
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", response.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	payments[index].Status = response.Status
	payments[index].ConfirmationStatus = response.ConfirmationStatus
}

func confirmPayment(index int) {
	response := services.PaymentConfirm(payments[index].PaymentId)
	fmt.Println("Payment Confirmation Response:")
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", response.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	payments[index].Status = response.Status
	payments[index].ConfirmationStatus = response.ConfirmationStatus
}

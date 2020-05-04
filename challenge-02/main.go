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
	VoidId             string
	CancellationStatus int
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
			if len(payments) > 0 {
				menuPayment(len(payments) - 1)
			}
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
		fmt.Println("  3 - Void Payment")
		fmt.Println("  4 - Select Void")
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
		case 3:
			voidPayment(index)
			if len(payments[index].Voids) > 0 {
				menuVoid(index, len(payments[index].Voids)-1)
			}
		case 4:
			if len(payments[index].Voids) > 0 {
				menuVoid(index, selectVoid(index))
			} else {
				fmt.Println("No previous voids for this payment to select")
			}
			selectVoid(index)
		case 0:
			break loopPayment
		default:
			fmt.Println(option, "is not a valid option")
		}
		fmt.Println()
	}
}

func selectVoid(index int) int {
	fmt.Println("Select a Void:")
	for i, v := range payments[index].Voids {
		fmt.Println(" ", i, "-", v.VoidId)
	}
	fmt.Println()

	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)
	fmt.Println()

	if option >= len(payments[index].Voids) {
		fmt.Println("This void doesn't exists. Selecting latest as default")
		option = len(payments[index].Voids) - 1
	}

	return option
}

func menuVoid(index int, voidIndex int) {
loopVoid:
	for {
		fmt.Println("Selected Payment:")
		fmt.Println("    PaymentId:", payments[index].PaymentId)
		fmt.Println("    Status:", payments[index].Status)
		fmt.Println("    ConfirmationStatus:", payments[index].ConfirmationStatus)
		fmt.Println()
		fmt.Println("Selected Void:")
		fmt.Println("    VoidId:", payments[index].Voids[voidIndex].VoidId)
		fmt.Println("    CancellationStatus:", payments[index].Voids[voidIndex].CancellationStatus)
		fmt.Println()
		fmt.Println("Please inform the operation you would like to execute:")
		fmt.Println("  1 - Reverse Void")
		fmt.Println("  0 - Go back")
		fmt.Println()

		var option int
		fmt.Print("Option: ")
		fmt.Scan(&option)
		fmt.Println()

		switch option {
		case 1:
			reverseVoid(index, voidIndex)
		case 0:
			break loopVoid
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

	success, response := services.Payment(amount)
	fmt.Print("Payment Response: ")
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure...")
	}
	fmt.Println("    PaymentId:", response.Payment.PaymentId)
	fmt.Println("    Status:", response.Payment.Status)
	fmt.Println("    ConfirmationStatus:", response.Payment.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.Payment.ReturnCode)
	fmt.Println("    ReturnMessage:", response.Payment.ReturnMessage)
	fmt.Println()

	if response.Payment.PaymentId != "" {
		newPayment := payment{
			PaymentId:          response.Payment.PaymentId,
			Status:             response.Payment.Status,
			ConfirmationStatus: response.Payment.ConfirmationStatus,
		}
		payments = append(payments, newPayment)
	}
}

func reversePayment(index int) {
	success, response := services.PaymentReverse(payments[index].PaymentId)
	fmt.Print("Payment Reversal Response: ")
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure...")
	}
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", response.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	if success {
		payments[index].Status = response.Status
		payments[index].ConfirmationStatus = response.ConfirmationStatus
	}
}

func confirmPayment(index int) {
	success, response := services.PaymentConfirm(payments[index].PaymentId)
	fmt.Print("Payment Confirmation Response: ")
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure...")
	}
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", response.ConfirmationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	if success {
		payments[index].Status = response.Status
		payments[index].ConfirmationStatus = response.ConfirmationStatus
	}
}

func voidPayment(index int) {
	success, response := services.Void(payments[index].PaymentId)
	fmt.Print("Payment Void Response: ")
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure...")
	}
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", response.ConfirmationStatus)
	fmt.Println("    VoidId:", response.VoidId)
	fmt.Println("    CancellationStatus:", response.CancellationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	if response.VoidId != "" {
		newVoid := void{
			VoidId:             response.VoidId,
			CancellationStatus: response.CancellationStatus,
		}
		payments[index].Voids = append(payments[index].Voids, newVoid)
		payments[index].Status = response.Status
	}
}

func reverseVoid(index int, voidIndex int) {
	success, response := services.VoidReverse(payments[index].PaymentId, payments[index].Voids[voidIndex].VoidId)
	fmt.Print("Void Reversal Response: ")
	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure...")
	}
	fmt.Println("    PaymentId:", payments[index].PaymentId)
	fmt.Println("    Status:", response.Status)
	fmt.Println("    ConfirmationStatus:", payments[index].ConfirmationStatus)
	fmt.Println("    VoidId:", payments[index].Voids[voidIndex].VoidId)
	fmt.Println("    CancellationStatus:", response.CancellationStatus)
	fmt.Println("    ReturnCode:", response.ReturnCode)
	fmt.Println("    ReturnMessage:", response.ReturnMessage)
	fmt.Println()

	if response.ReturnCode == 0 {
		payments[index].Status = response.Status
		payments[index].ConfirmationStatus = response.ConfirmationStatus
	}
}

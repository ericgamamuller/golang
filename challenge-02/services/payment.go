package services

import (
	"git/challenge-02/clients"
	"git/challenge-02/contracts"
	"git/challenge-02/services/builders"
)

func Payment(amount int) (bool, contracts.PaymentResponse) {
	response := clients.PaymentPost(builders.PaymentRequest(amount), GetToken())
	success := (response.Payment.ReturnCode == "000")

	return success, response
}

func PaymentReverse(paymentId string) (bool, contracts.PaymentReversalResponse) {
	response := clients.PaymentReversalDelete(paymentId, builders.PaymentReversalRequest(), GetToken())
	success := (response.ReasonCode == 0)

	return success, response
}

func PaymentConfirm(paymentId string) (bool, contracts.PaymentConfirmationResponse) {
	response := clients.PaymentConfirmationPut(paymentId, builders.PaymentConfirmationRequest(), GetToken())
	success := (response.ReasonCode == 0)

	return success, response
}

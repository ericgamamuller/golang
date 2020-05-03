package services

import (
	"git/challenge-02/clients"
	"git/challenge-02/contracts"
	"git/challenge-02/services/builders"
)

func Payment(amount int) contracts.PaymentResponse {
	return clients.PaymentPost(builders.PaymentRequest(amount), GetToken())
}

func PaymentReverse(paymentId string) contracts.PaymentReversalResponse {
	return clients.PaymentReversalDelete(paymentId, builders.PaymentReversalRequest(), GetToken())
}

func PaymentConfirm(paymentId string) contracts.PaymentConfirmationResponse {
	return clients.PaymentConfirmationPut(paymentId, builders.PaymentConfirmationRequest(), GetToken())
}

package services

import (
	"git/challenge-02/clients"
	"git/challenge-02/contracts"
	"git/challenge-02/services/builders"
)

func Void(paymentId string) (bool, contracts.VoidResponse) {
	response := clients.VoidPost(paymentId, builders.VoidRequest(), GetToken())
	success := response.ReasonCode == 0

	return success, response
}

func VoidReverse(paymentId string, voidId string) (bool, contracts.VoidReversalResponse) {
	response := clients.VoidReversalDelete(paymentId, voidId, builders.VoidReversalRequest(), GetToken())
	success := response.ReasonCode == 0

	return success, response
}

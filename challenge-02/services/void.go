package services

import (
	"git/challenge-02/contracts"
)

func Void(request contracts.VoidRequest) contracts.VoidResponse {
	return contracts.VoidResponse{}
}

func VoidReverse(request contracts.VoidReversalRequest) contracts.VoidReversalResponse {
	return contracts.VoidReversalResponse{}
}

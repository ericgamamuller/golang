package builders

import "git/challenge-02/contracts"

func VoidRequest() contracts.VoidRequest {
	return contracts.VoidRequest{
		MerchantVoidId:   "1587997030607",
		MerchantVoidDate: "2020-04-27T14:21:37.176Z",
		Card: contracts.VoidCard{
			InputMode:  "Typed",
			CardNumber: "5432123454321234",
		},
	}
}

func VoidReversalRequest() contracts.VoidReversalRequest {
	return contracts.VoidReversalRequest{}
}

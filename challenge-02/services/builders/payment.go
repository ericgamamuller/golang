package builders

import "git/challenge-02/contracts"

func PaymentRequest(amount int) contracts.PaymentRequest {
	return contracts.PaymentRequest{
		MerchantOrderId: "1587997030607",
		Payment: contracts.Payment{
			Type:            "PhysicalCreditCard",
			SoftDescriptor:  "Desafio GO 2",
			PaymentDateTime: "2020-01-08T11:00:00",
			Amount:          amount,
			Installmentes:   1,
			Interest:        "ByMerchant",
			Capture:         true,
			ProductId:       1,
			CreditCard: contracts.CreditCard{
				CardNumber:                     "5432123454321234",
				ExpirationDate:                 "12/2021",
				SecurityCodeStatus:             "Collected",
				SecurityCode:                   "123",
				BrandId:                        1,
				IssuerId:                       401,
				InputMode:                      "Typed",
				AuthenticationMethod:           "NoPassword",
				TruncateCardNumberWhenPrinting: true,
			},
			PinPadInformation: contracts.PinPad{
				PhysicalCharacteristics: "PinPadWithChipReaderWithoutSamAndContactless",
				ReturnDataInfo:          "00",
				SerialNumber:            "0820471929",
				TerminalId:              "42004558",
			},
		},
	}
}

func PaymentReversalRequest() contracts.PaymentReversalRequest {
	return contracts.PaymentReversalRequest{}
}

func PaymentConfirmationRequest() contracts.PaymentConfirmationRequest {
	return contracts.PaymentConfirmationRequest{}
}

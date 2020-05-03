package clients

import (
	"encoding/json"
	"fmt"
	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

func PaymentConfirmationPut(paymentId string, request contracts.PaymentConfirmationRequest, token contracts.AuthToken) contracts.PaymentConfirmationResponse {
	url := config.ReadString("Payment.BaseUrl") + "/1/physicalSales/" + paymentId + "/confirmation"

	body, _ := json.Marshal(request)

	var response contracts.PaymentConfirmationResponse

	client := resty.New()
	_, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		SetResult(&response).
		Put(url)

	if err != nil {
		fmt.Println("Failed to Confirm Payment! Error:", err)
	}

	return response
}

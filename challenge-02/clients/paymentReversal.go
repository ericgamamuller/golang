package clients

import (
	"encoding/json"
	"fmt"

	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

func PaymentReversalDelete(paymentId string, request contracts.PaymentReversalRequest, token contracts.AuthToken) contracts.PaymentReversalResponse {
	url := config.ReadString("Payment.BaseUrl") + "/1/physicalSales/" + paymentId

	body, _ := json.Marshal(request)

	var response contracts.PaymentReversalResponse

	client := resty.New()
	_, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		SetResult(&response).
		Delete(url)

	if err != nil {
		fmt.Println("Failed to Reverse Payment! Error:", err)
	}

	return response
}

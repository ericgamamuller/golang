package clients

import (
	"encoding/json"
	"fmt"

	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

func VoidReversalDelete(paymentId string, voidId string, request contracts.VoidReversalRequest, token contracts.AuthToken) contracts.VoidReversalResponse {
	url := config.ReadString("Payment.BaseUrl") + "/1/physicalSales/" + paymentId + "/voids/" + voidId

	body, _ := json.Marshal(request)

	var response contracts.VoidReversalResponse

	client := resty.New()
	_, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		SetResult(&response).
		Delete(url)

	if err != nil {
		fmt.Println("Failed to Reverse Void! Error:", err)
	}

	return response
}

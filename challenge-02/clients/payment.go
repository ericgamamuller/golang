package clients

import (
	"encoding/json"
	"fmt"
	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

func PaymentPost(request contracts.PaymentRequest, token contracts.AuthToken) contracts.PaymentResponse {
	url := config.ReadString("Payment.BaseUrl") + "/1/physicalSales"

	body, _ := json.Marshal(request)

	var response contracts.PaymentResponse

	client := resty.New()
	_, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to Post Payment! Error:", err)
	}

	return response
}

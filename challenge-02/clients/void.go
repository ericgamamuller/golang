package clients

import (
	"encoding/json"
	"fmt"
	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

func VoidPost(paymentId string, request contracts.VoidRequest, token contracts.AuthToken) contracts.VoidResponse {
	url := config.ReadString("Payment.BaseUrl") + "/1/physicalSales/" + paymentId + "/voids"

	body, _ := json.Marshal(request)

	var response contracts.VoidResponse

	client := resty.New()
	_, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to Post Void! Error:", err)
	}

	return response
}

package clients

import (
	"fmt"

	"git/challenge-03/config"
	"git/challenge-03/contracts"

	resty "github.com/go-resty/resty/v2"
)

const url = "https://authsandbox.braspag.com.br"

func TokenRequest(request contracts.AuthRequest) (int, contracts.AuthToken) {
	url := config.ReadString("Auth.BaseUrl")

	var response contracts.AuthToken

	client := resty.New()
	result, err := client.R().
		SetBasicAuth(request.Username, request.Password).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody("grant_type=client_credentials").
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to get Auth response! Error:", err)
	}

	return result.RawResponse.StatusCode, response
}

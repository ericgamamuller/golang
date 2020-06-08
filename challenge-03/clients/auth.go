package clients

import (
	"fmt"

	"git/challenge-02/config"
	"git/challenge-02/contracts"

	resty "github.com/go-resty/resty/v2"
)

const url = "https://authsandbox.braspag.com.br"

func TokenRequest(request contracts.AuthRequest) contracts.AuthToken {
	url := config.ReadString("Auth.BaseUrl") + "/oauth2/token"

	var response contracts.AuthToken

	client := resty.New()
	_, err := client.R().
		SetBasicAuth(request.Username, request.Password).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody("grant_type=client_credentials").
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to get Auth response! Error:", err)
	}

	return response
}

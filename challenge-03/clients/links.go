package clients

import (
	"fmt"
	"git/challenge-03/config"
	"git/challenge-03/contracts"

	resty "github.com/go-resty/resty/v2"
)

func CreateLink(request contracts.CreateLinkRequest, token contracts.AuthToken) (int, contracts.CreateLinkResponse) {
	url := config.ReadString("Api.BaseUrl")

	var response contracts.CreateLinkResponse

	client := resty.New()
	result, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to create payment link! Error:", err)
	}

	if result.RawResponse.StatusCode != 201 {
		fmt.Println("Response error for status", result.RawResponse.StatusCode, ":", result)
	}

	return result.RawResponse.StatusCode, response
}

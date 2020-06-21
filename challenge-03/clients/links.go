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
	ret, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(url)

	status := ret.StatusCode()

	if err != nil {
		fmt.Println("Failed to create payment link! Error:", err)
	} else if status != 201 {
		fmt.Println("Response error for status", status, "with return:", ret)
	}

	return status, response
}

func GetLink(id string, token contracts.AuthToken) (int, contracts.CreateLinkResponse) {
	url := config.ReadString("Api.BaseUrl") + id

	var response contracts.CreateLinkResponse

	client := resty.New()
	ret, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		Get(url)

	status := ret.StatusCode()

	if err != nil {
		fmt.Println("Failed to get payment link! Error:", err)
	} else if status != 200 {
		fmt.Println("Response error for status", status, "with return:", ret)
	}

	return status, response
}

func UpdateLink(id string, request contracts.CreateLinkRequest, token contracts.AuthToken) (int, contracts.CreateLinkResponse) {
	url := config.ReadString("Api.BaseUrl") + id

	var response contracts.CreateLinkResponse

	client := resty.New()
	ret, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Put(url)

	status := ret.StatusCode()

	if err != nil {
		fmt.Println("Failed to update payment link! Error:", err)
	} else if status != 200 {
		fmt.Println("Response error for status", status, "with return:", ret)
	}

	return status, response
}

func DeleteLink(id string, token contracts.AuthToken) int {
	url := config.ReadString("Api.BaseUrl") + id

	client := resty.New()
	ret, err := client.R().
		SetAuthToken(token.AccessToken).
		SetHeader("Content-Type", "application/json").
		Delete(url)

	status := ret.StatusCode()

	if err != nil {
		fmt.Println("Failed to get payment link! Error:", err)
	} else if status != 200 && status != 204 {
		fmt.Println("Response error for status", status, "with return:", ret)
	}

	return status
}

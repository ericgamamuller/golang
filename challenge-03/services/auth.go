package services

import (
	"time"

	"git/challenge-03/clients"
	"git/challenge-03/config"
	"git/challenge-03/contracts"
)

var token contracts.AuthToken

func GetToken() contracts.AuthToken {
	if !tokenIsValid() {
		request := contracts.AuthRequest{
			config.ReadString("Auth.ClientId"),
			config.ReadString("Auth.ClientSecret")}

		token = clients.TokenRequest(request)

		token.ExpiresAt = time.Now().Second() + token.ExpiresIn
	}

	return token
}

func tokenIsValid() bool {
	if token.AccessToken == "" {
		return false
	}

	if token.ExpiresAt <= time.Now().Second() {
		return false
	}

	return true
}

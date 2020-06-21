package services

import (
	"strconv"
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

		var status int
		status, token = clients.TokenRequest(request)

		if status != 201 {
			panic("Authentication failed!")
		}

		expiresIn, err := strconv.Atoi(token.ExpiresIn)
		if err != nil {
			panic(err.Error())
		}

		token.ExpiresAt = time.Now().Second() + expiresIn - 30
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

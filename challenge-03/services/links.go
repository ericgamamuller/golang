package services

import (
	"fmt"
	"git/challenge-03/clients"
	"git/challenge-03/contracts"
)

func CreateLink(request contracts.CreateLinkRequest) (bool, contracts.CreateLinkResponse) {
	status, response := clients.CreateLink(request, GetToken())

	if status == 201 {
		return true, response
	}

	fmt.Println("Create link failed with status:", status)

	return false, response
}

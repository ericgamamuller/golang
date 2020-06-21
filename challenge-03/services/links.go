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

func GetLink(id string) (bool, contracts.CreateLinkResponse) {
	status, response := clients.GetLink(id, GetToken())

	if status == 200 {
		return true, response
	}

	fmt.Println("Get link failed with status:", status)

	return false, response
}

func UpdateLink(id string, request contracts.CreateLinkRequest) (bool, contracts.CreateLinkResponse) {
	status, response := clients.UpdateLink(id, request, GetToken())

	if status == 200 {
		return true, response
	}

	fmt.Println("Update link failed with status:", status)

	return false, response
}

func DeleteLink(id string) bool {
	status := clients.DeleteLink(id, GetToken())

	if status == 200 || status == 204 {
		return true
	}

	fmt.Println("Delete link failed with status:", status)

	return false
}

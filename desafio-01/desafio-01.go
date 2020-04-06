package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	healtcheck("https://parametersdownloadsandbox.cieloecommerce.cielo.com.br/healthcheck")
}

type healtcheckResponse struct {
	IsHealthy bool
}

func healtcheck(url string) bool {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error calling '"+url+"':", err)
		return false
	}

	if response.StatusCode != 200 {
		return false
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading '"+url+"' response:", err)
		return false
	}

	var bodyObject healtcheckResponse
	json.Unmarshal(body, &bodyObject)

	return bodyObject.IsHealthy
}

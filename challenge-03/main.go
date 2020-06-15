package main

import (
	"fmt"
	"net/http"

	"git/challenge-03/config"
	"git/challenge-03/routes"
)

const version = "0.1"

func main() {
	config.Initialize()
	routes.Load()
	fmt.Println("Escutando porta 8000")
	http.ListenAndServe(":8000", nil)
}

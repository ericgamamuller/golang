package main

import (
	"git/challenge-03/config"
	"net/http"
)

const version = "0.1"

func main() {
	config.Initialize()
	routes.Load()
	http.ListenAndServe(":8000", nil)
}

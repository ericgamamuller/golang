package routes

import (
	"net/http"

	"git/challenge-03/controllers"
)

func Load() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/create", controllers.Create)
}

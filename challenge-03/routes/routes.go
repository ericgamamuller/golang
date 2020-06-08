package routes

import "net/http"

func Load() {
	http.HandleFunc("/", controllers.LinksIndex)
}

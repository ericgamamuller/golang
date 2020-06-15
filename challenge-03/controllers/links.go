package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"git/challenge-03/contracts"
	"git/challenge-03/models"
	"git/challenge-03/services"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	links := models.GetLinks()
	templates.ExecuteTemplate(w, "LinksIndex", links)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "LinksNew", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		request := contracts.CreateLinkRequest{}

		request.Type = r.FormValue("type")
		request.Name = r.FormValue("name")
		request.Description = r.FormValue("description")
		request.ShowDescription = true
		tempPrice, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			panic(err.Error())
		}
		request.Price = tempPrice
		request.Weight = 100
		request.ExpirationDate = "2021-06-15"
		tempQuantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			panic(err.Error())
		}
		request.Quantity = tempQuantity
		request.SKU = "teste"

		request.Shipping = contracts.Shipping{}
		request.Shipping.Type = "WithoutShipping"

		success, response := services.CreateLink(request)

		if success {
			link := models.PaymentLink{
				ID:       response.ID,
				ShortURL: response.ShortURL,
				Name:     response.Name,
				Type:     response.Type,
				Price:    response.Price,
			}
			link.InsertLink()
		}
	}
	http.Redirect(w, r, "/", 301)
}

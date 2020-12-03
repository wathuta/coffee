package handlers

import (
	"net/http"

	"github.com/wathuta/rest-go/data"
	"github.com/wathuta/rest-go/model"
)

//swagger:route POST /create product createProduct
//Creates a product in the database
//responses:
//	501:internalServerError
//	201:created

//CreateProduct is a handler function that is responsible for handling the post requests
func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(model.Product)
	data.AddProduct(&prod)
}

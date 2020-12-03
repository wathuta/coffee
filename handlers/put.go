package handlers

import (
	"net/http"

	"github.com/wathuta/rest-go/data"
	"github.com/wathuta/rest-go/model"
)

//swagger:route PUT /update/{id} product updateProduct
//makes changes to a product in the database.The product is identified by the id
//	responses:
//		200:created
//		501:internalServerError
//		201: noContent

//UpdateProduct is a handler function that is responsible for handling the update requests
func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := getID(r)

	prod := r.Context().Value(KeyProduct{}).(model.Product)
	err := data.UpdateProduct(id, &prod)
	if err != nil {
		p.l.Fatal(err)
	}
}

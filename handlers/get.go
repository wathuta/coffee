package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wathuta/rest-go/data"
)

// swagger:route GET /get products listproducts
// Returns a list of products
// Responses:
//		200:getResponse

//GetProduct is a handler function that is responsible for handling the get requests.it returns products from the data store
func (p *Product) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	prod := data.GetAllProduct()
	err := ToJSON(w, prod)
	if err != nil {
		p.l.Fatal("unable to marshal json", err)
	}
}

//swagger:route GET /get/{id} product listSingleProduct
//Returns a single product
// Responses:
//		200:created

//GetSingleProduct is a handler funvtion that is responsible for handling the get requests.
//It returns  a single product based on the id
func (p *Product) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := getID(r)
	prod := data.GetSingleProduct(id)
	err := json.NewEncoder(w).Encode(prod)
	if err != nil {
		p.l.Fatal("unable to marshal json", err)
	}
}

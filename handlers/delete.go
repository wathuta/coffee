package handlers

import (
	"fmt"
	"net/http"

	"github.com/wathuta/rest-go/data"
)

// swagger:route DELETE /delete/{id} products deleteProduct
// Deletes a product from the database
// Responses:
//		201:noContent

//DeleteProduct is a handler function that is responsible for handling the delete requests
//it deltes a product from the database
func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getID(r)
	fmt.Println(id)
	err := data.DeleteProduct(id)
	if err == data.ErrorNotFound {
		http.Error(w, "the product you want to delete does not exist ", http.StatusInternalServerError)
	} else if err != nil {
		p.l.Fatal(err)
	}
}

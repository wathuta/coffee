//Package classification of Product-API
//
//Documentation for Product API
//
//	Scheme:http
//	BasePath:/
//	Version:1.0.0
//	Host:localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//	swagger:meta
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/wathuta/rest-go/model"

	"github.com/gorilla/mux"
)

//Product implements the handler interface
type Product struct {
	l *log.Logger
}

//GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

//NewProduct is the entry point to the handler
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

//ToJSON is responsible for serialising information to json data
func ToJSON(w http.ResponseWriter, p []*model.Product) error {
	return json.NewEncoder(w).Encode(p)
}

func getID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("unable to find id", err)
	}
	return id
}

package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wathuta/rest-go/model"
)

//KeyProduct is used as a key value for the object in the request context
type KeyProduct struct{}

//MiddlewareFunc implements the handler interface.
//middlewares are a good accessory to use while authenticating or doing cors with js
//htey can intercept requests and carry out authentication
func (p *Product) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var prod model.Product
		err := json.NewDecoder(r.Body).Decode(&prod)
		if err != nil {
			http.Error(w, "unable to unmarshal json", http.StatusInternalServerError)
			p.l.Fatal(err)
			return
		}
		//Validation of the input
		err = prod.Validate()
		if err != nil {
			http.Error(w, "Error validating product", http.StatusInternalServerError)
			p.l.Fatal("[ERROR] validating product", err)
			return
		}

		//add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

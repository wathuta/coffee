package handlers

import (
	"github.com/wathuta/rest-go/model"
)

//A list of all the products in the database product for the Get method
//swagger:response getResponse
type productsResponse struct {
	//in:body
	Body []*model.Product
}

//error response for invalid data format or internal server Error
//swagger:response internalServerError
type serverErrorWrapper struct {
	//in:body
	Error GenericError
}

//Response for a delete function
//swagger:response noContent
type noContentWrapper struct {
}

//Response for  creation of an object
//swagger:response created
type createdWrapper struct {
	//in:body
	Body model.Product
}

//swagger:parameters createProduct updateProduct
type productParamWrapper struct {
	//Product data structure to create or a update a product
	//in:body
	//required:true
	Body model.Product
}

//swagger:parameters deleteProduct updateProduct listSingleProduct
type productIDParamWrapper struct {
	//the id of the product
	//in:path
	//required:true
	ID int `json:"id"`
}

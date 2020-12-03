package model

import (
	"regexp"

	"github.com/go-playground/validator"
)

//Product is a data model of a product
//swagger:model
type Product struct {
	//the id of the product
	//required:true
	//min:1
	ID int `json:"id"`
	//name of the  product
	//required :true
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	SKU         string `json:"sku" validate:"required,sku"`
	Price       int    `json:"price" validate:"gte=0,required"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}

//Validate method does validation by tags using govalidator
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}
func validateSku(fl validator.FieldLevel) bool {
	//sku abc-cds-sf format.checking the input with the expected format
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	//the sku can only be one for a product
	if len(matches) != 1 {
		return false
	}
	return true
}

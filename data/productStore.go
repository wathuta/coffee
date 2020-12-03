package data

import (
	"net/http"
	"time"

	"github.com/wathuta/rest-go/model"
)

//GetAllProduct returns a slice of alll the products inthe store
func GetAllProduct() []*model.Product {
	return productList
}

//GetSingleProduct  returns  a  product from the database
func GetSingleProduct(id int) *model.Product {
	for j, i := range productList {
		if productList[j].ID == id {
			return i
		}
	}
	return nil
}

//AddProduct appends a product to the products store
func AddProduct(p *model.Product) int {
	p.ID = getNextID()
	productList = append(productList, p)

	return http.StatusOK
}

func getNextID() int {
	return len(productList) + 1
}

//ErrorNotFound is a custom made error for unfruitful searches
var ErrorNotFound error

//UpdateProduct alters the information about a product stored in the database
func UpdateProduct(id int, p *model.Product) error {
	for j, i := range productList {
		if i.ID == id {
			productList[j] = p
			productList[j].ID = id
			return nil
		}
	}
	return ErrorNotFound
}

//DeleteProduct removes entities from the storage
func DeleteProduct(id int) error {
	i := findIndexbyID(id)
	if i == -1 {
		return ErrorNotFound
	}
	productList = append(productList[:i], productList[i+1])
	return nil
}

//finds index of an object in the database system and returns the -1 when product cannot be found
func findIndexbyID(id int) int {
	for j, i := range productList {
		if i.ID == id {
			return j
		}
	}
	return -1
}

var productList = []*model.Product{
	&model.Product{
		ID:          1,
		Name:        "latte",
		Description: "Frothy milky coffee",
		Price:       2,
		SKU:         "as-d-sdf",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&model.Product{
		ID:          2,
		Name:        "Esspreso",
		Description: "Frothy milky coffee",
		Price:       354,
		SKU:         "irk-jmn-vof",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

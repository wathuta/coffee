package model

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "brian",
		Price: 1,
		SKU:   "acv-we-ad",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

package dto

import (
	repository "github.com/ThailanTec/go-hexagonal/application/repository/product"
)

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *repository.DProduct) (*repository.DProduct, error) {
	if p.ID != "" {
		product.Product.ID = p.ID
	}
	product.Product.Name = p.Name
	product.Product.Price = p.Price
	product.Product.Status = p.Status

	_, err := product.IsValid()

	if err != nil {
		return &repository.DProduct{}, err
	}
	return product, nil
}

package domain

import (
	uuid "github.com/satori/go.uuid"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float, optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		Status: DISABLED,
		ID:     uuid.NewV4().String(),
	}
	return &product
}

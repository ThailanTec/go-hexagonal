package repository

import (
	"errors"

	"github.com/ThailanTec/go-hexagonal/application/core/domain"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type DProduct struct {
	Product *domain.Product
}

func NewProduct() *DProduct {
	product := &DProduct{&domain.Product{Status: domain.DISABLED, ID: uuid.NewV4().String()}}
	return product
}

func (p *DProduct) IsValid() (bool, error) {

	if p.Product.Status == "" {
		p.Product.Status = domain.DISABLED
	}

	if p.Product.Status != domain.ENABLED && p.Product.Status != domain.DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Product.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *DProduct) Enable() error {
	if p.Product.Price > 0 {
		p.Product.Status = domain.ENABLED
		return nil
	}
	return errors.New("the price most be greater than zero to enable product")
}

func (p *DProduct) Disable() error {
	if p.Product.Price == 0 {
		p.Product.Status = domain.DISABLED
		return nil
	}
	return errors.New("the price must be zero in order to have the product disabled")
}

func (p *DProduct) GetID() string {
	return p.Product.Status
}

func (p *DProduct) GetName() string {
	return p.Product.Name
}

func (p *DProduct) GetPrice() float64 {

	return p.Product.Price

}

func (p *DProduct) GetStatus() string {
	return p.Product.Status
}

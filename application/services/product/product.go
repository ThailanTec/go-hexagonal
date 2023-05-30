package services

import (
	"github.com/ThailanTec/go-hexagonal/application/core/ports"
	repository "github.com/ThailanTec/go-hexagonal/application/repository/product"
)

type ProductService struct {
	Persistence ports.ProductPesistenceInterface
}

func NewProductService(persistence ports.ProductPesistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ports.ProductInterface, error) {

	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Reconfigurar ação newProduct e implementar teste unitario
func (s *ProductService) Create(name string, price float64) (ports.ProductInterface, error) {
	product := repository.NewProduct()

	product.Product.Name = name
	product.Product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return product, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}
	return result, nil
}

func (s *ProductService) Enabled(product ports.ProductInterface) (ports.ProductInterface, error) {

	err := product.Enable()
	if err != nil {
		return product, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}

	return result, nil
}

func (s *ProductService) Disable(product ports.ProductInterface) (ports.ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return product, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return product, err
	}

	return result, nil
}

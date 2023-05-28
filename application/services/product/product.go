package services

import "github.com/ThailanTec/go-hexagonal/application/core/ports"

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
func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()

	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

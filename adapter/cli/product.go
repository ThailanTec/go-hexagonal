package cli

import (
	"fmt"

	"github.com/ThailanTec/go-hexagonal/application/core/ports"
)

func Run(service ports.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with then name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", res.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with then name %s has been created with the price %f and status %s", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil

}

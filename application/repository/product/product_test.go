package repository_test

import (
	"testing"

	"github.com/ThailanTec/go-hexagonal/application/core/domain"
	repository "github.com/ThailanTec/go-hexagonal/application/repository/product"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	product := repository.NewProduct()

	product.Product.Name = "Hello"
	product.Product.ID = uuid.NewV4().String()
	product.Product.Status = domain.ENABLED
	product.Product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Product.Status = "INVALID"
	_, err = product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Product.Status = domain.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

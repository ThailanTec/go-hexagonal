package application

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{}

	product.Name = "Liquidificador"
	product.ID = "1"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = -10
	err = product.Enable()
	require.Equal(t, "the price most be greater than zero to enable product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := Product{}

	product.Name = "Liquidificador"
	product.ID = "1"
	product.Status = ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{}

	product.Name = "Hello"
	product.ID = uuid.NewV4().String()
	product.Status = ENABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

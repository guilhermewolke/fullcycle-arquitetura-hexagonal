package application_test

import (
	"github.com/guilhermewoelke/arquitetura-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "O preço do produto precisa ser maior que zero", err.Error())
}

func TestApplicationProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "O preço precisa ser zero para que ele possa ser desativado", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestApplicationProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "O produto precisa estar \"enabled\" ou \"disabled\"", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "O preço precisa ser maior ou igual a zero", err.Error())

	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)
}

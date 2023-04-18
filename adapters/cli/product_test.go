package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/guilhermewoelke/arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/guilhermewoelke/arquitetura-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product test"
	productPrice := 25.99
	productStatus := "enabled"
	productID := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID '%s' with the name '%s' has been created with the price '%.2f' and status '%s'!", productID, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID '%s' (%s) has been enabled!", productID, productName)
	result, err = cli.Run(service, "enable", productID, "", float64(0))
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID '%s' (%s) has been disabled!", productID, productName)
	result, err = cli.Run(service, "disable", productID, "", float64(0))
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s | Name: %s | Price: %.2f | Status: %s", productID, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productID, "", float64(0))
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

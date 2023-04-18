package cli

import (
	"fmt"

	"github.com/guilhermewoelke/arquitetura-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action, productID string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID '%s' with the name '%s' has been created with the price '%.2f' and status '%s'!", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID '%s' (%s) has been enabled!", res.GetID(), res.GetName())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID '%s' (%s) has been disabled!", res.GetID(), res.GetName())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s | Name: %s | Price: %.2f | Status: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil

}

package catalog

import (
	"checkout_system/models"
	"checkout_system/upload"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func LoadPrices() (models.Products, error) {
	products := make(map[string]models.Product)

	productList, err := upload.ReadJSONCatalog()
	if err != nil {
		return models.Products{}, errors.Wrap(err, "cannot load catalog")
	}

	for _, product := range productList {
		products[product.SKU] = models.Product{SKU: product.SKU, Name: product.Name, Price: product.Price}
	}
	logrus.Info("Catalog loaded.")

	return models.Products{Products: products}, nil
}

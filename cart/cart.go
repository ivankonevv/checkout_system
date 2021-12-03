package cart

import (
	"checkout_system/catalog"
	"checkout_system/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

type ProductCart struct {
	Products map[string][]models.Product
}

var AvailableProducts models.Products

func LoadCatalog() {
	products, err := catalog.LoadPrices()
	if err != nil {
		logrus.Error("an error occurred in LoadPrices():", err)
		return
	}
	AvailableProducts = products

	return
}

func New() *ProductCart {
	return &ProductCart{Products: make(map[string][]models.Product)}
}

func (c *ProductCart) Contains(sku string) bool {
	_, ok := c.Products[sku]
	if !ok {
		return false
	}

	return true
}

func (c *ProductCart) Get(sku string) ProductCart {
	products := make(map[string][]models.Product)
	product, ok := c.Products[sku]
	if !ok {
		return ProductCart{}
	}
	products[sku] = product

	return ProductCart{Products: products}

}

func (c ProductCart) Len(sku string) int {
	return len(c.Products[sku])
}

func (c *ProductCart) UpdateItemsPrice(sku string, price float32) {
	for i, _ := range c.Products[sku] {
		c.Products[sku][i] = models.Product{
			SKU:   sku,
			Name:  c.Products[sku][i].Name,
			Price: price,
		}
	}

	return
}

func (c *ProductCart) UpdateOneItemPrice(sku string, price float32) {
	products := c.Products[sku]
	for i, p := range products {
		if p.Price == price {
			continue
		}
		c.Products[sku][i] = models.Product{SKU: sku, Name: c.Products[sku][i].Name, Price: price}
		return
	}

	return
}

func (c *ProductCart) Scan(sku string) {
	c.Products[sku] = append(c.Products[sku], models.Product{
		SKU:   sku,
		Name:  AvailableProducts.Products[sku].Name,
		Price: AvailableProducts.Products[sku].Price,
	})

	return
}

func (c *ProductCart) Total() float32 {
	var totalAmount float32
	for sku, products := range c.Products {
		var productAmount float32
		for _, product := range products {
			totalAmount += product.Price
			productAmount += product.Price
		}
		logrus.WithFields(logrus.Fields{
			"SKU":   sku,
			"count": len(products),
			"total": fmt.Sprintf("%.2f", productAmount),
		}).Info(products[0].Name)
	}
	logrus.Infof("Total Amount: $%.2f", totalAmount)

	return totalAmount
}

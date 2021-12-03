package main

import (
	"checkout_system/cart"
	"checkout_system/offers"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	if err := godotenv.Load(".env"); err != nil {
		logrus.Error("an error occurred while load env variables:", err)
		return
	}
	return
}

func main() {
	cart.LoadCatalog()
	myShoppingCart := cart.New()

	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")

	myShoppingCart.Scan("atv")
	myShoppingCart.Scan("atv")
	myShoppingCart.Scan("atv")

	myShoppingCart.Scan("mbp")
	myShoppingCart.Scan("mbp")

	myShoppingCart.Scan("vga")
	myShoppingCart.Scan("vga")

	specialOffers := []func(productCart *cart.ProductCart) *cart.ProductCart{
		offers.ApplyIpadPrice,
		offers.ApplyAppleTVSpecialPrice,
		offers.ApplyVGASuperDeal,
	}

	for _, o := range specialOffers {
		myShoppingCart = o(myShoppingCart)
	}

	myShoppingCart.Total()
}
